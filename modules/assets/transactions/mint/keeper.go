package mint

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper mapper.Mapper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	mutables := types.NewMutables(message.Properties, message.MaintainersID)
	immutables := types.NewImmutables(message.Properties)
	assetID := mapper.NewAssetID(message.ChainID, message.MaintainersID, message.ClassificationID, immutables.GetHashID())
	asset := mapper.NewAsset(assetID, mutables, immutables, message.Lock, message.Burn)
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	if assets.Get(assetID) != nil {
		return constants.EntityAlreadyExists
	}
	assets.Add(asset)
	return nil
}

func NewTransactionKeeper(Mapper types.Mapper) types.TransactionKeeper {
	switch value := Mapper.(type) {
	case mapper.Mapper:
		return transactionKeeper{mapper: value}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization, module %v", constants.ModuleName)))
	}
}
