package burn

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
)

type Keeper interface {
	transact(sdkTypes.Context, message) error
}

type baseKeeper struct {
	mapper mapper.Mapper
}

func NewKeeper(mapper mapper.Mapper) Keeper {
	return baseKeeper{mapper: mapper}
}

var _ Keeper = (*baseKeeper)(nil)

func (baseKeeper baseKeeper) transact(context sdkTypes.Context, message message) error {
	assets := baseKeeper.mapper.Assets(context, message.assetID)
	asset := assets.Asset(message.assetID)
	if asset == nil {
		return constants.EntityNotFoundCode
	}
	return assets.Remove(asset)
}
