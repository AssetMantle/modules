package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper mapper.Mapper
}

func NewTransactionKeeper(mapper mapper.Mapper) types.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := msg.(Message)
	immutablePropertyList := message.Properties.GetList()
	hashID := transactionKeeper.mapper.MakeHashID(immutablePropertyList)
	assetID := mapper.NewAssetID(message.ChainID, message.MaintainersID, message.ClassificationID, hashID)
	asset := mapper.NewAsset(assetID, message.Properties, message.Lock, message.Burn)
	assets := transactionKeeper.mapper.Assets(context, assetID)
	if assets.Get(assetID) != nil {
		return constants.EntityAlreadyExistsCode
	}
	assets.Add(asset)
	return nil
}
