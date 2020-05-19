package send

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
	asset, Error := baseKeeper.mapper.Read(context, mapper.NewAssetAddress(message.Address))
	if Error != nil {
		return Error
	}
	asset.SetOwner(message.To)
	return baseKeeper.mapper.Update(context, asset)
}
