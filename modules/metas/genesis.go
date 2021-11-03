package metas

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/keeper"
	metaTypes "github.com/persistenceOne/persistenceSDK/modules/metas/types"
	abci "github.com/tendermint/tendermint/abci/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdkTypes.Context, keeper keeper.Keeper, data *metaTypes.GenesisState) (res []abci.ValidatorUpdate) {
	keeper.SetParameters(ctx, data.Parameters)
	for _, meta := range data.Mappables {
		keeper.SetMeta(ctx, meta)
	}
	return res
}

func ExportGenesis(ctx sdkTypes.Context, keeper keeper.Keeper) *metaTypes.GenesisState {
	return &metaTypes.GenesisState{
		Parameters: keeper.GetParameters(ctx),
		Mappables:  keeper.GetAllMetas(ctx),
	}
}

func ValidateGenesis(data *metaTypes.GenesisState) error {
	return data.Parameters.Validate()
}
