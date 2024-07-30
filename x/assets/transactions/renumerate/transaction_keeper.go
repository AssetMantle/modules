// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/properties"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	authorizeAuxiliary    helpers.Auxiliary
	renumerateAuxiliary   helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if !transactionKeeper.parameterManager.Fetch(context).GetParameter(propertyConstants.RenumerateEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("renumerate is not enabled")
	}

	fromAddress := message.GetFromAddress()

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}

	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.AssetID))

	Mappable := assets.GetMappable(key.NewKey(message.AssetID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("asset with ID %s not found", message.AssetID.AsString())
	}
	asset := mappable.GetAsset(Mappable)

	if _, err := transactionKeeper.authorizeAuxiliary.GetKeeper().Help(context, authorize.NewAuxiliaryRequest(asset.GetClassificationID(), message.FromID, constants.CanRenumerateAssetPermission)); err != nil {
		return nil, err
	}

	supply := asset.GetSupply()

	// TODO: Clarify why supply property can't be meta,
	// The first if statement is to check if the supply property is not meta and the second if statement is to check if the supply property is meta
	if supplyProperty := asset.GetProperty(propertyConstants.SupplyProperty.GetID()); supplyProperty != nil && !supplyProperty.IsMeta() {
		auxiliaryResponse, err := transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(supplyProperty))
		if err != nil {
			return nil, err
		}

		if supplyProperty = supplement.GetMetaPropertiesFromResponse(auxiliaryResponse).GetProperty(propertyConstants.SupplyProperty.GetID()); supplyProperty != nil && supplyProperty.IsMeta() {
			//TODO: This type casting will panic if the property is not a meta property
			supply = supplyProperty.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
		} else {
			return nil, errorConstants.MetaDataError.Wrapf("supply property is not revealed")
		}
	}

	if supply.IsNegative() {
		return nil, errorConstants.MetaDataError.Wrapf("asset supply is negative")
	}

	if _, err := transactionKeeper.renumerateAuxiliary.GetKeeper().Help(context, renumerate.NewAuxiliaryRequest(message.FromID, message.AssetID, supply)); err != nil {
		return nil, err
	}

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
			case renumerate.Auxiliary.GetName():
				transactionKeeper.renumerateAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
