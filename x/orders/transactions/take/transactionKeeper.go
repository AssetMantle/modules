// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package take

import (
	"context"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/burn"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/mappable"
	"github.com/AssetMantle/modules/x/orders/module"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
	"github.com/AssetMantle/schema/go/properties/utilities"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	burnAuxiliary         helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	address, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(address, message.FromID)); err != nil {
		return nil, err
	}

	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.OrderID))

	Mappable := orders.Get(key.NewKey(message.OrderID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("order with ID %s not found", message.OrderID.AsString())
	}
	order := mappable.GetOrder(Mappable)

	if order.GetTakerID().Compare(baseIDs.PrototypeIdentityID()) != 0 && order.GetTakerID().Compare(message.FromID) != 0 {
		return nil, errorConstants.NotAuthorized.Wrapf("taker ID %s is not authorized to take private order with ID %s", message.FromID.AsString(), message.OrderID.AsString())
	}
	takerOwnableSplit, err := sdkTypes.NewDecFromStr(message.TakerOwnableSplit)
	makerReceiveTakerOwnableSplit := order.GetMakerOwnableSplit().ToDec().MulTruncate(order.GetExchangeRate()).MulTruncate(sdkTypes.SmallestDec())
	takerReceiveMakerOwnableSplit := takerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(order.GetExchangeRate())

	switch updatedMakerOwnableSplit := order.GetMakerOwnableSplit().Sub(takerReceiveMakerOwnableSplit.TruncateInt()); {
	case updatedMakerOwnableSplit.Equal(sdkTypes.ZeroInt()):
		if takerOwnableSplit.LT(makerReceiveTakerOwnableSplit) {
			return nil, errorConstants.InsufficientBalance.Wrapf("taker ownable split %s is less than the required amount %s for order execution", message.TakerOwnableSplit, makerReceiveTakerOwnableSplit.String())
		}

		orders.Remove(mappable.NewMappable(order))
	case updatedMakerOwnableSplit.LT(sdkTypes.ZeroInt()):
		if takerOwnableSplit.LT(makerReceiveTakerOwnableSplit) {
			return nil, errorConstants.InsufficientBalance.Wrapf("taker ownable split %s is less than the required amount %s for order execution", message.TakerOwnableSplit, makerReceiveTakerOwnableSplit.String())
		}

		takerReceiveMakerOwnableSplit = order.GetMakerOwnableSplit().ToDec()

		orders.Remove(mappable.NewMappable(order))
	default:
		makerReceiveTakerOwnableSplit = takerOwnableSplit
		mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewNumberData(updatedMakerOwnableSplit)))

		orders.Mutate(mappable.NewMappable(base.NewOrder(order.GetClassificationID(), order.GetImmutables(), order.GetMutables().Mutate(utilities.AnyPropertyListToPropertyList(mutableProperties.GetList()...)...))))
	}

	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, order.GetMakerID(), order.GetTakerOwnableID(), makerReceiveTakerOwnableSplit.TruncateInt())); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, message.FromID, order.GetMakerOwnableID(), takerReceiveMakerOwnableSplit.TruncateInt())); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.burnAuxiliary.GetKeeper().Help(context, burn.NewAuxiliaryRequest(order.GetClassificationID())); err != nil {
		return nil, err
	}

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
				transactionKeeper.authenticateAuxiliary = value
			case burn.Auxiliary.GetName():
				transactionKeeper.burnAuxiliary = value
			}
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
