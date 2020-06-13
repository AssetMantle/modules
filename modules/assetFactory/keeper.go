package assetFactory

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mint"
)

type Keeper interface {
	getMintKeeper() mint.Keeper

	getAssetQuerier() asset.Querier
}

type keeper struct {
	mintKeeper mint.Keeper

	assetQuerier asset.Querier
}

func NewKeeper(codec *codec.Codec, storeKey sdkTypes.StoreKey, paramSpace params.Subspace) Keeper {
	Mapper := mapper.NewMapper(codec, storeKey)
	return keeper{
		mintKeeper:   mint.NewKeeper(Mapper),
		assetQuerier: asset.NewQuerier(Mapper),
	}
}

var _ Keeper = (*keeper)(nil)

func (keeper keeper) getMintKeeper() mint.Keeper { return keeper.mintKeeper }

func (keeper keeper) getAssetQuerier() asset.Querier { return keeper.assetQuerier }
