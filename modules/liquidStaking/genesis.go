/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package liquidStaking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	keeper.SetParams(ctx, data.Params)
}

func ExportGenesis(ctx sdk.Context, keeper Keeper) *GenesisState {
	params := keeper.GetParams(ctx)
	return NewGenesisState(params)
}
