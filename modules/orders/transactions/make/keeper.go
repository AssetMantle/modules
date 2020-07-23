package make

import (
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type transactionKeeper struct {
	mapper utilities.Mapper
}

var _ utilities.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	message.Salt = base.NewHeight(context.BlockHeight())

	orderHash := message.GenerateHash()

	makerSignature := base.NewSignature(base.NewID("makerAddress"), message.From.Bytes(), base.NewHeight(context.BlockHeight()))
	orderHashProperty := base.NewProperty(base.NewID("orderHash"), base.NewFact(orderHash.String(), base.NewSignatures([]types.Signature{makerSignature})))
	properties := message.Properties.Add(orderHashProperty)
	mutables := base.NewMutables(properties, message.MaintainersID)
	immutables := base.NewImmutables(properties)
	orderID := mapper.NewOrderID(base.NewID(context.ChainID()), mutables.GetMaintainersID(), message.ClassificationID, immutables.GetHashID())
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

func initializeTransactionKeeper(mapper utilities.Mapper, externalKeepers []interface{}) utilities.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
