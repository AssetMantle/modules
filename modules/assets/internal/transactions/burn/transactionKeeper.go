// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"context"
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/unbond"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/schema/data"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	maintainAuxiliary     helpers.Auxiliary
	renumerateAuxiliary   helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	unbondAuxiliary       helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	if !transactionKeeper.parameterManager.Fetch(context).GetParameter(constants.BurnEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("burning is not enabled")
	}

	fromAddress, _ := sdkTypes.AccAddressFromBech32(message.From)

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}
	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.AssetID))

	Mappable := assets.Get(key.NewKey(message.AssetID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("asset with ID %s not found", message.AssetID.AsString())
	}
	asset := mappable.GetAsset(Mappable)

	auxiliaryResponse, err := transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(asset.GetBurn()))
	if err != nil {
		return nil, err
	}
	metaProperties := supplement.GetMetaPropertiesFromResponse(auxiliaryResponse)

	if burnHeightMetaProperty := metaProperties.GetProperty(constants.BurnHeightProperty.GetID()); burnHeightMetaProperty != nil {
		burnHeight := burnHeightMetaProperty.Get().(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
		if burnHeight.Compare(baseTypes.NewHeight(sdkTypes.UnwrapSDKContext(context).BlockHeight())) > 0 {
			return nil, errorConstants.NotAuthorized.Wrapf("burning is not allowed until height %d", burnHeight.Get())
		}
	}

	if _, err := transactionKeeper.renumerateAuxiliary.GetKeeper().Help(context, renumerate.NewAuxiliaryRequest(message.FromID, message.AssetID, sdkTypes.ZeroDec())); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.unbondAuxiliary.GetKeeper().Help(context, unbond.NewAuxiliaryRequest(asset.GetClassificationID(), fromAddress)); err != nil {
		return nil, err
	}

	assets.Remove(mappable.NewMappable(asset))

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
			case maintain.Auxiliary.GetName():
				transactionKeeper.maintainAuxiliary = value
			case renumerate.Auxiliary.GetName():
				transactionKeeper.renumerateAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case unbond.Auxiliary.GetName():
				transactionKeeper.unbondAuxiliary = value
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
