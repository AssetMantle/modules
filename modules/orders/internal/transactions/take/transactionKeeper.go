// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package take

import (
	"context"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/cosmos/cosmos-sdk/types"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameters            helpers.Parameters
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
}

func (transactionKeeper transactionKeeper) Take(ctx context.Context, message *Message) (*TransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context types.Context, message helpers.Message) helpers.TransactionResponse {
	_, err := transactionKeeper.Handle(context.Context(), message.(*Message))
	return newTransactionResponse(err)
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*Response, error) {

	address, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), authenticate.NewAuxiliaryRequest(address, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, errorConstants.EntityNotFound
	}

	orders := transactionKeeper.mapper.NewCollection(types.UnwrapSDKContext(context)).Fetch(key.NewKey(message.OrderID))

	Mutable := orders.Get(key.NewKey(message.OrderID))
	if Mutable == nil {
		return nil, errorConstants.EntityNotFound
	}
	order := Mutable.(documents.Order)

	if order.GetTakerID().Compare(baseIDs.PrototypeIdentityID()) != 0 && order.GetTakerID().Compare(message.FromID) != 0 {
		return nil, errorConstants.NotAuthorized
	}

	makerReceiveTakerOwnableSplit := order.GetMakerOwnableSplit().MulTruncate(order.GetExchangeRate()).MulTruncate(types.SmallestDec())
	takerReceiveMakerOwnableSplit := message.TakerOwnableSplit.QuoTruncate(types.SmallestDec()).QuoTruncate(order.GetExchangeRate())

	switch updatedMakerOwnableSplit := order.GetMakerOwnableSplit().Sub(takerReceiveMakerOwnableSplit); {
	case updatedMakerOwnableSplit.Equal(types.ZeroDec()):
		if message.TakerOwnableSplit.LT(makerReceiveTakerOwnableSplit) {
			return nil, errorConstants.InsufficientBalance
		}

		orders.Remove(mappable.NewMappable(order))
	case updatedMakerOwnableSplit.LT(types.ZeroDec()):
		if message.TakerOwnableSplit.LT(makerReceiveTakerOwnableSplit) {
			return nil, errorConstants.InsufficientBalance
		}

		takerReceiveMakerOwnableSplit = order.GetMakerOwnableSplit()

		orders.Remove(mappable.NewMappable(order))
	default:
		makerReceiveTakerOwnableSplit = message.TakerOwnableSplit
		mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(updatedMakerOwnableSplit)))

		orders.Mutate(mappable.NewMappable(base.NewOrder(order.GetClassificationID(), order.GetImmutables(), order.GetMutables().Mutate(mutableProperties.GetList()...))))
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(message.FromID, order.GetMakerID(), order.GetTakerOwnableID(), makerReceiveTakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(module.ModuleIdentityID, message.FromID, order.GetMakerOwnableID(), takerReceiveMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	return &Response{}, nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

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
			}
		default:
			panic(errorConstants.UninitializedUsage)
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
