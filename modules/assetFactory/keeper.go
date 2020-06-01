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

type keeper struct {
	burnKeeper   burn.Keeper
	mintKeeper   mint.Keeper
	mutateKeeper mutate.Keeper
	assetQuerier asset.Querier
}

func NewKeeper(codec *codec.Codec, storeKey sdkTypes.StoreKey, paramSpace params.Subspace) Keeper {
	Mapper := mapper.NewMapper(codec, storeKey)
	return keeper{
		burnKeeper:   burn.NewKeeper(Mapper),
		mintKeeper:   mint.NewKeeper(Mapper),
		mutateKeeper: mutate.NewKeeper(Mapper),
		assetQuerier: asset.NewQuerier(Mapper),
	}
}

var _ Keeper = (*keeper)(nil)

func (keeper keeper) getBurnKeeper() burn.Keeper     { return keeper.burnKeeper }
func (keeper keeper) getMintKeeper() mint.Keeper     { return keeper.mintKeeper }
func (keeper keeper) getMutateKeeper() mutate.Keeper { return keeper.mutateKeeper }
func (keeper keeper) getAssetQuerier() asset.Querier { return keeper.assetQuerier }
