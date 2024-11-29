// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/properties"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
	baseTypes "github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/unbond"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/purge"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	authorizeAuxiliary    helpers.Auxiliary
	purgeAuxiliary        helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	unbondAuxiliary       helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	if !transactionKeeper.parameterManager.Fetch(context).Get().GetParameter(propertyConstants.BurnEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("burning is not enabled")
	}

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message)); err != nil {
		return nil, err
	}

	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.AssetID))

	Mappable := assets.GetMappable(key.NewKey(message.AssetID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("asset with ID %s not found", message.AssetID.AsString())
	}
	asset := mappable.GetAsset(Mappable)

	if _, err := transactionKeeper.authorizeAuxiliary.GetKeeper().Help(context, authorize.NewAuxiliaryRequest(asset.GetClassificationID(), message.GetFromIdentityID(), constants.CanBurnAssetPermission)); err != nil {
		return nil, err
	}

	burnHeight := asset.GetBurnHeight()

	if burnHeightProperty := asset.GetProperty(propertyConstants.BurnHeightProperty.GetID()); burnHeightProperty != nil && !burnHeightProperty.IsMeta() {
		auxiliaryResponse, err := transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(burnHeightProperty))
		if err != nil {
			return nil, err
		}

		if burnHeightProperty = supplement.GetMetaPropertiesFromResponse(auxiliaryResponse).GetProperty(propertyConstants.BurnHeightProperty.GetID()); burnHeightProperty != nil && burnHeightProperty.IsMeta() {
			burnHeight = burnHeightProperty.Get().(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
		} else {
			return nil, errorConstants.MetaDataError.Wrapf("burn height property is not revealed")
		}
	}

	if burnHeight.Compare(baseTypes.CurrentHeight(context)) > 0 {
		return nil, errorConstants.NotAuthorized.Wrapf("burning is not allowed until height %d", burnHeight.Get())
	}

	supply := asset.GetSupply()

	if supplyProperty := asset.GetProperty(propertyConstants.SupplyProperty.GetID()); supplyProperty != nil && !supplyProperty.IsMeta() {
		auxiliaryResponse, err := transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(supplyProperty))
		if err != nil {
			return nil, err
		}

		if supplyMetaProperty := supplement.GetMetaPropertiesFromResponse(auxiliaryResponse).GetProperty(propertyConstants.SupplyProperty.GetID()); supplyMetaProperty != nil && supplyMetaProperty.IsMeta() {
			supply = supplyMetaProperty.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
		} else {
			return nil, errorConstants.MetaDataError.Wrapf("assets without revealed supply cannot be burned")
		}
	}

	if _, err := transactionKeeper.purgeAuxiliary.GetKeeper().Help(context, purge.NewAuxiliaryRequest(message.GetFromIdentityID(), message.AssetID, supply)); err != nil {
		return nil, err
	}

	bondAmount := sdkTypes.ZeroInt()
	if bondAmountProperty := asset.GetProperty(propertyConstants.BondAmountProperty.GetID()); bondAmountProperty == nil || !bondAmountProperty.IsMeta() {
		return nil, errorConstants.MetaDataError.Wrapf("asset with ID %s has no revealed bond amount", message.AssetID)
	} else {
		bondAmount = bondAmountProperty.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}

	if _, err := transactionKeeper.unbondAuxiliary.GetKeeper().Help(context, unbond.NewAuxiliaryRequest(asset.GetClassificationID(), message.GetFromAddress(), bondAmount)); err != nil {
		return nil, err
	}

	assets.Remove(record.NewRecord(asset))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper
	transactionKeeper.parameterManager = parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case authorize.Auxiliary.GetName():
				transactionKeeper.authorizeAuxiliary = value
			case purge.Auxiliary.GetName():
				transactionKeeper.purgeAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case unbond.Auxiliary.GetName():
				transactionKeeper.unbondAuxiliary = value
			}
		}
	}

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
