package cancel

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(message.OrderID)
	order := orders.Get(message.OrderID)
	if order == nil {
		return constants.EntityNotFound
	}
	//check if from address is provisioned in makerID
	//if !order.GetMakerID().Equals(message.From) {
	//	return constants.NotAuthorized
	//}
	if !order.CanBurn(base.NewHeight(context.BlockHeight())) {
		return constants.DeletionNotAllowed
	}
	orders.Remove(order)
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, _ []interface{}) helpers.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
