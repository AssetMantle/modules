/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/persistenceOne/persistenceSDK/modules/liquidStaking/types"
)

// GetBlockHeight randomized BlockHeight
func GetBlockHeight(r *rand.Rand) uint64 {
	return uint64(r.Intn(1000))
}

// RandomizedGenState generates a random GenesisState for liquidStaking
func RandomizedGenState(simState *module.SimulationState) {

	// params
	blocksPerYear := uint64(2 * 60 * 60 * 8766 / 5)
	liquidStakingGenesis := types.NewGenesisState(types.NewParams(blocksPerYear))

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(liquidStakingGenesis)
}
