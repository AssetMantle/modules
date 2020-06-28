package assetFactory

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/types"
)

type Keeper interface {
	getMintKeeper() types.TransactionKeeper

	getAssetQuerier() types.QueryKeeper
}

type keeper struct {
	mintKeeper types.TransactionKeeper

	assetQuerier types.QueryKeeper
}

func NewKeeper(codec *codec.Codec, storeKey sdkTypes.StoreKey, paramSpace params.Subspace) Keeper {
	Mapper := mapper.NewMapper(codec, storeKey)
	return keeper{
		mintKeeper:   mint.NewTransactionKeeper(Mapper),
		assetQuerier: asset.NewQueryKeeper(Mapper),
	}
}

var _ Keeper = (*keeper)(nil)

func (keeper keeper) getMintKeeper() types.TransactionKeeper { return keeper.mintKeeper }

func (keeper keeper) getAssetQuerier() types.QueryKeeper { return keeper.assetQuerier }
