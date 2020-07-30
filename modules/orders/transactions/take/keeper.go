package take

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/auxiliaries/swap"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionKeeper struct {
	mapper         helpers.Mapper
	exchangeKeeper helpers.AuxiliaryKeeper
	bankKeeper     bank.Keeper
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

	order = mapper.NewOrder(order.GetID(), order.GetBurn(), order.GetLock(), order.GetImmutables(),
		order.GetMakerAddress(), message.From, order.GetMakerAssetAmount(), order.GetMakerAssetData(), order.GetTakerAssetAmount(),
		order.GetTakerAssetData(), order.GetSalt())
	orders = orders.Mutate(order)
	if Error := transactionKeeper.exchangeKeeper.Help(context, swap.NewAuxiliaryRequest(order, transactionKeeper.bankKeeper)); Error != nil {
		return Error
	}
	orders.Remove(order)
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, externalKeepers []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case bank.Keeper:
			transactionKeeper.bankKeeper = value
		case helpers.AuxiliaryKeeper:
			transactionKeeper.exchangeKeeper = value
		}
	}
	return transactionKeeper
}
