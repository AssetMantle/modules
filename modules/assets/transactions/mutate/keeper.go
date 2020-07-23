package mutate

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type transactionKeeper struct {
	mapper utilities.Mapper `json:"mapper" valid:"required~Enter the Mapper"`
}

var _ utilities.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	assetID := message.AssetID
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	asset := assets.Get(assetID)
	if asset == nil {
		return constants.EntityNotFound
	}
	mutableProperties := asset.GetMutables().Get()
	for _, property := range message.Properties.GetList() {
		if mutableProperties.Get(property.GetID()) == nil {
			return constants.EntityNotFound
		}
		mutableProperties = mutableProperties.Mutate(property)
	}
	asset = mapper.NewAsset(asset.GetID(), asset.GetBurn(), asset.GetLock(), asset.GetImmutables(), base.NewMutables(mutableProperties, asset.GetMutables().GetMaintainersID()))
	assets = assets.Mutate(asset)
	return nil
}

func initializeTransactionKeeper(mapper utilities.Mapper, _ []interface{}) utilities.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
