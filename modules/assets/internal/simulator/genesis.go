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
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters/dummy"
	assetsModule "github.com/persistenceOne/persistenceSDK/modules/assets/module"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"math/rand"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {

	var data types.Data
	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		dummy.ID.String(),
		&data,
		simulationState.Rand,
		func(rand *rand.Rand) { data = base.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	// TODO add assetList
	genesisState := genesis.Prototype().Initialize(nil, []types.Parameter{dummy.Parameter.Mutate(data)})

	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", codec.MustMarshalJSONIndent(simulationState.Cdc, genesisState))
	simulationState.GenState[assetsModule.Name] = simulationState.Cdc.MustMarshalJSON(genesisState)
}
