package mutate

import (
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
	message := messageFromInterface(msg)
	assetID := message.AssetID

	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	asset := assets.Get(assetID)

	if asset == nil {
		return constants.EntityNotFound
	}
	mutableProperties := asset.GetMutables().Get()

	assetStruct, _ := asset.(mapper.Order)
	transactionKeeper.bankKeeper.SendCoinsFromModuleToAccount(context, constants.ModuleName, message.From, sdkTypes.Coins{assetStruct.SellOrder})
	transactionKeeper.bankKeeper.SendCoins(context, message.From, assetStruct.From, sdkTypes.Coins{assetStruct.BuyOrder})

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
