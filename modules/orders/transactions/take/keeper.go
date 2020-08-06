/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package take

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/auxiliaries/swap"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionKeeper struct {
	mapper            helpers.Mapper
	exchangeAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	orderID := message.OrderID
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(orderID)
	order := orders.Get(orderID)
	if order == nil {
		return constants.EntityNotFound
	}
	if order.GetTakerAddress() != nil && !message.From.Equals(order.GetTakerAddress()) {
		return constants.NotAuthorized
	}
	order = mapper.NewOrder(order.GetID(), order.GetBurn(), order.GetLock(), order.GetImmutables(),
		order.GetMakerAddress(), message.From, order.GetMakerAssetAmount(), order.GetMakerAssetData(), order.GetTakerAssetAmount(),
		order.GetTakerAssetData(), order.GetSalt())
	orders = orders.Mutate(order)
	if Error := transactionKeeper.exchangeAuxiliary.GetKeeper().Help(context, swap.NewAuxiliaryRequest(order)); Error != nil {
		return Error
	}
	orders.Remove(order)
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case swap.Auxiliary.GetName():
				transactionKeeper.exchangeAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
