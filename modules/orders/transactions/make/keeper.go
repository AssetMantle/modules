package make

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
	message.Salt = base.NewHeight(context.BlockHeight())
	orderHash := message.GenerateHash()
	orderHashProperty := base.NewProperty(base.NewID(OrderHash), base.NewFact(orderHash.String(), nil))
	properties := message.Properties.Add(orderHashProperty)
	immutables := base.NewImmutables(properties)
	orderID := mapper.NewOrderID(base.NewID(context.ChainID()), immutables.GetHashID())
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(orderID)
	if orders.Get(orderID) != nil {
		return constants.EntityAlreadyExists
	}
	orders.Add(mapper.NewOrder(orderID, message.Burn, message.Lock, immutables, message.From, message.TakerAddress,
		message.MakerAssetAmount, message.MakerAssetData, message.TakerAssetAmount, message.TakerAssetData, message.Salt))
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, _ []interface{}) helpers.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
