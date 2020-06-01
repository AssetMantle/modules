package mutate

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type Keeper interface {
	transact(sdkTypes.Context, Message) error
}

type keeper struct {
	mapper mapper.Mapper
}

func NewKeeper(mapper mapper.Mapper) Keeper {
	return keeper{mapper: mapper}
}

var _ Keeper = (*keeper)(nil)

func (keeper keeper) transact(context sdkTypes.Context, message Message) error {
	immutablePropertyList := message.propertyList
	hashID := keeper.mapper.GenerateHashID(immutablePropertyList)
	assetID := keeper.mapper.GenerateAssetID(message.chainID, message.maintainersID, message.classificationID, hashID)
	asset := keeper.mapper.MakeAsset(assetID, &types.BaseProperties{PropertyList: message.propertyList}, message.lock, message.burn)
	assets := keeper.mapper.Assets(context, assetID)
	if assets.Get(assetID) == nil {
		return constants.EntityNotFoundCode
	}
	return assets.Mutate(asset)
}
