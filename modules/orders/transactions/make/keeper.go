package make

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper     helpers.Mapper
	BankKeeper bank.Keeper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	salt := base.NewHeight(context.BlockHeight())
	orderHash := message.GenerateHash(salt)
	orderHashProperty := base.NewProperty(base.NewID(OrderHash), base.NewFact(orderHash.String(), nil))
	properties := message.Properties.Add(orderHashProperty)
	immutables := base.NewImmutables(properties)
	orderID := mapper.NewOrderID(base.NewID(context.ChainID()), immutables.GetHashID())
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(orderID)
	if orders.Get(orderID) != nil {
		return constants.EntityAlreadyExists
	}

	orders.Add(mapper.NewOrder(orderID, message.Burn, message.Lock, immutables, message.FromID, message.ToID,
		message.MakerAssetAmount, message.MakerAssetData, message.TakerAssetAmount, message.TakerAssetData, salt))
	return nil
}

func checkAndWrapCoin(assetType types.ID, assetData types.ID, assetAmount sdkTypes.Dec) (types.ID, error) {
	// check if id is in assets/ splits, if not wrap coin.
	return nil, nil
}

func takeCustody(context sdkTypes.Context, bankKeeper bank.Keeper, makerAddress sdkTypes.AccAddress, asset types.ID) error {
	//convert to split, put in orderdb
	// if coin -> send to module, if split -> create identity, send to module identity
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, externalKeepers []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case bank.Keeper:
			transactionKeeper.BankKeeper = value
		}
	}
	return transactionKeeper
}
