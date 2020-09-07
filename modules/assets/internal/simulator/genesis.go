/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters/dummy"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"math/rand"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {

	var data types.Data
	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		dummy.Key,
		&data,
		simulationState.Rand,
		func(rand *rand.Rand) { data = base.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	Parameters := baseHelpers.NewParameters(dummy.Parameter.Mutate(data))

	// TODO add assetList
	genesisState := genesis.NewGenesisState(nil, Parameters)

	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", codec.MustMarshalJSONIndent(simulationState.Cdc, genesisState))
	simulationState.GenState[mapper.ModuleName] = simulationState.Cdc.MustMarshalJSON(genesisState)
}
