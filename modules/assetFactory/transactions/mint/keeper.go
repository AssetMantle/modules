package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
)

type Keeper interface {
	transact(sdkTypes.Context, Message) error
}

type baseKeeper struct {
	mapper mapper.Mapper
}

func NewKeeper(mapper mapper.Mapper) Keeper {
	return baseKeeper{mapper: mapper}
}

var _ Keeper = (*baseKeeper)(nil)

func (baseKeeper baseKeeper) transact(context sdkTypes.Context, message Message) error {
	return baseKeeper.mapper.Create(context, mapper.NewAssetAddress(message.Address), message.To, message.Lock)
}
