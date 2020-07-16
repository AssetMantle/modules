package mutate

import (
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper     types.Mapper
	bankKeeper bank.Keeper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	fmt.Println("hello")
	message := messageFromInterface(msg)
	assetID := message.AssetID
	fmt.Println("hello2")

	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	asset := assets.Get(assetID)
	fmt.Println("hell3")

	if asset == nil {
		return constants.EntityNotFound
	}
	fmt.Println("hell4")

	mutableProperties := asset.GetMutables().Get()
	fmt.Println("hell5")

	//for _, property := range message.Properties.GetList() {
	//	if mutableProperties.Get(property.GetID()) == nil {
	//		return constants.MutableNotFound
	//	}
	//	mutableProperties = mutableProperties.Mutate(property)
	//}
	assetStruct, ok := asset.(mapper.Order)
	fmt.Println(ok)
	transactionKeeper.bankKeeper.SubtractCoins(context, assetStruct.From, sdkTypes.Coins{assetStruct.SellOrder})
	transactionKeeper.bankKeeper.AddCoins(context, assetStruct.From, sdkTypes.Coins{assetStruct.BuyOrder})
	transactionKeeper.bankKeeper.SubtractCoins(context, message.From, sdkTypes.Coins{assetStruct.BuyOrder})
	transactionKeeper.bankKeeper.AddCoins(context, message.From, sdkTypes.Coins{assetStruct.SellOrder})
	asset = mapper.NewOrder(asset.GetID(), message.From, assetStruct.BuyOrder, assetStruct.SellOrder, types.NewMutables(mutableProperties, asset.GetMutables().GetMaintainersID()), asset.GetImmutables(), types.NewHeight(5), asset.GetBurn())
	assets = assets.Mutate(asset)
	return nil
}

func initializeTransactionKeeper(mapper types.Mapper, externalKeepers []interface{}) types.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case bank.Keeper:
			transactionKeeper.bankKeeper = value
		}
	}
	return transactionKeeper
}
