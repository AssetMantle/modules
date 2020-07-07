package assets

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assets/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mutate"
	"github.com/persistenceOne/persistenceSDK/types"
)

type Keeper interface {
	getAssetQuerier() types.QueryKeeper

	getBurnKeeper() types.TransactionKeeper
	getMintKeeper() types.TransactionKeeper
	getMutateKeeper() types.TransactionKeeper
}

type keeper struct {
	AssetQuerier types.QueryKeeper

	BurnKeeper   types.TransactionKeeper
	MintKeeper   types.TransactionKeeper
	MutateKeeper types.TransactionKeeper
}

func NewKeeper(codec *codec.Codec, storeKey sdkTypes.StoreKey, paramSpace params.Subspace) Keeper {
	Mapper := mapper.NewMapper(codec, storeKey)
	return keeper{
		AssetQuerier: asset.NewQueryKeeper(Mapper),
		BurnKeeper:   burn.NewTransactionKeeper(Mapper),
		MintKeeper:   mint.NewTransactionKeeper(Mapper),
		MutateKeeper: mutate.NewTransactionKeeper(Mapper),
	}
}

var _ Keeper = (*keeper)(nil)

func (keeper keeper) getAssetQuerier() types.QueryKeeper { return keeper.AssetQuerier }

func (keeper keeper) getBurnKeeper() types.TransactionKeeper   { return keeper.BurnKeeper }
func (keeper keeper) getMintKeeper() types.TransactionKeeper   { return keeper.MintKeeper }
func (keeper keeper) getMutateKeeper() types.TransactionKeeper { return keeper.MutateKeeper }
