package mint

import (
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper types.Mapper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	message.Salt = types.NewHeight(context.BlockHeight())

	orderHash := message.GenerateHash()

	makerSignature := types.NewSignature(types.NewID("makerAddress"), message.From.Bytes(), types.NewHeight(context.BlockHeight()))
	orderHashProperty := types.NewProperty(types.NewID("orderHash"), types.NewFact(orderHash.String(), types.NewSignatures([]types.Signature{makerSignature})))
	properties := message.Properties.Add(orderHashProperty)
	mutables := types.NewMutables(properties, message.MaintainersID)
	immutables := types.NewImmutables(properties)
	orderID := mapper.NewOrderID(types.NewID(context.ChainID()), mutables.GetMaintainersID(), message.ClassificationID, immutables.GetHashID())
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(orderID)
	if orders.Get(orderID) != nil {
		return constants.EntityAlreadyExists
	}
	fmt.Println("1")
	orders.Add(mapper.NewOrder(orderID, message.Burn, message.Lock, immutables, mutables, message.From, message.TakerAddress,
		message.MakerAssetAmount, message.MakerAssetData, message.TakerAssetAmount, message.TakerAssetData, message.Salt))
	fmt.Println("2")
	return nil
}

func initializeTransactionKeeper(mapper types.Mapper, externalKeepers []interface{}) types.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
