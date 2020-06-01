package assetFactory

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mutate"
)

type Keeper interface {
	getBurnKeeper() burn.Keeper
	getMintKeeper() mint.Keeper
	getMutateKeeper() mutate.Keeper
	getAssetQuerier() asset.Querier
}

type baseKeeper struct {
	burnKeeper   burn.Keeper
	mintKeeper   mint.Keeper
	mutateKeeper mutate.Keeper
	assetQuerier asset.Querier
}

func NewKeeper(codec *codec.Codec, storeKey sdkTypes.StoreKey, paramSpace params.Subspace) Keeper {
	Mapper := mapper.NewMapper(codec, storeKey)
	return baseKeeper{
		burnKeeper:   burn.NewKeeper(Mapper),
		mintKeeper:   mint.NewKeeper(Mapper),
		mutateKeeper: mutate.NewKeeper(Mapper),
		assetQuerier: asset.NewQuerier(Mapper),
	}
}

var _ Keeper = (*baseKeeper)(nil)

func (baseKeeper baseKeeper) getBurnKeeper() burn.Keeper     { return baseKeeper.burnKeeper }
func (baseKeeper baseKeeper) getMintKeeper() mint.Keeper     { return baseKeeper.mintKeeper }
func (baseKeeper baseKeeper) getMutateKeeper() mutate.Keeper { return baseKeeper.mutateKeeper }
func (baseKeeper baseKeeper) getAssetQuerier() asset.Querier { return baseKeeper.assetQuerier }
