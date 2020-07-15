package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper types.Mapper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	mutables := types.NewMutables(message.Properties, message.MaintainersID)
	immutables := types.NewImmutables(message.Properties)
	assetID := mapper.NewAssetID(types.NewID(context.ChainID()), mutables.GetMaintainersID(), message.ClassificationID, immutables.GetHashID())
	asset := mapper.NewAsset(assetID, message.Burn, message.Lock, immutables, mutables)
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	if assets.Get(assetID) != nil {
		return constants.EntityAlreadyExists
	}
	assets.Add(asset)
	return nil
}

func initializeTransactionKeeper(mapper types.Mapper, externalKeepers ...interface{}) types.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
