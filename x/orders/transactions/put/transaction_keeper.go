// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/documents/base"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
	baseTypes "github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/record"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if !transactionKeeper.parameterManager.Fetch(context).GetParameter(propertyConstants.PutEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("put orders not enabled")
	}

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message)); err != nil {
		return nil, err
	}

	makerSplit, ok := sdkTypes.NewIntFromString(message.MakerSplit)
	if !ok || makerSplit.IsNegative() {
		return nil, errorConstants.IncorrectFormat.Wrapf("maker split is not a valid integer")
	}

	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.GetFromIdentityID(), constants.ModuleIdentity.GetModuleIdentityID(), message.MakerAssetID, makerSplit)); err != nil {
		return nil, err
	}

	takerSplit, ok := sdkTypes.NewIntFromString(message.TakerSplit)
	if !ok || takerSplit.IsNegative() {
		return nil, errorConstants.IncorrectFormat.Wrapf("taker split is not a valid integer")
	}

	if message.ExpiryHeight.Compare(baseTypes.CurrentHeight(context)) <= 0 {
		return nil, errorConstants.InvalidRequest.Wrapf("order expiry is in the past")
	} else if message.ExpiryHeight.Get()-baseTypes.CurrentHeight(context).Get() > transactionKeeper.parameterManager.Fetch(context).GetParameter(propertyConstants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get() {
		return nil, errorConstants.InvalidRequest.Wrapf("order expiry exceeds maximum allowed %d", transactionKeeper.parameterManager.Fetch(context).GetParameter(propertyConstants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get())
	}

	putOrder := base.NewPutOrder(message.GetFromIdentityID(), message.MakerAssetID, message.TakerAssetID, makerSplit, takerSplit, message.ExpiryHeight)

	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(putOrder.GetPutOrderID()))
	if orders.GetMappable(key.NewKey(putOrder.GetPutOrderID())) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("order with ID %s already exists", putOrder.GetPutOrderID().AsString())
	}

	if err := putOrder.ValidateBasic(); err != nil {
		return nil, err
	}

	orders.Add(record.NewRecord(putOrder))

	return newTransactionResponse(putOrder.GetPutOrderID()), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, externalKeeper := range auxiliaries {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			}
		}
	}

	helpers.PanicOnUninitializedTransactionKeeperFields(transactionKeeper)
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
