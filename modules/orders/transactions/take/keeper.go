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
	mapper            helpers.Mapper
	exchangeAuxiliary helpers.Auxiliary
	bankKeeper        bank.Keeper
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
	if Error := transactionKeeper.exchangeAuxiliary.GetKeeper().Help(context, swap.NewAuxiliaryRequest(order, transactionKeeper.bankKeeper)); Error != nil {
		return Error
	}
	orders.Remove(order)
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case bank.Keeper:
			transactionKeeper.bankKeeper = value
		case helpers.Auxiliary:
			switch value.GetName() {
			case swap.Auxiliary.GetName():
				transactionKeeper.exchangeAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
