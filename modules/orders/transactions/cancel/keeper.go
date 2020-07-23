package cancel

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type transactionKeeper struct {
	mapper utility.Mapper
}

var _ utility.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(message.OrderID)
	order := orders.Get(message.OrderID)
	if order == nil {
		return constants.EntityNotFound
	}
	if !order.GetMakerAddress().Equals(message.From) {
		return constants.NotAuthorized
	}
	if !order.CanBurn(schema.NewHeight(context.BlockHeight())) {
		return constants.DeletionNotAllowed
	}
	orders.Remove(order)
	return nil
}

func initializeTransactionKeeper(mapper utility.Mapper, externalKeepers []interface{}) utility.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
