package make

import (
	"fmt"
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
	message.Salt = schema.NewHeight(context.BlockHeight())

	orderHash := message.GenerateHash()

	makerSignature := schema.NewSignature(schema.NewID("makerAddress"), message.From.Bytes(), schema.NewHeight(context.BlockHeight()))
	orderHashProperty := schema.NewProperty(schema.NewID("orderHash"), schema.NewFact(orderHash.String(), schema.NewSignatures([]schema.Signature{makerSignature})))
	properties := message.Properties.Add(orderHashProperty)
	mutables := schema.NewMutables(properties, message.MaintainersID)
	immutables := schema.NewImmutables(properties)
	orderID := mapper.NewOrderID(schema.NewID(context.ChainID()), mutables.GetMaintainersID(), message.ClassificationID, immutables.GetHashID())
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

func initializeTransactionKeeper(mapper utility.Mapper, externalKeepers []interface{}) utility.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
