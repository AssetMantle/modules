/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"fmt"
	"math/rand"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseSimulation "github.com/persistenceOne/persistenceSDK/utilities/simulation/schema/types/base"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/mappable"
	assetsModule "github.com/persistenceOne/persistenceSDK/modules/classifications/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/parameters/dummy"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
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

	mappableList := make([]helpers.Mappable, simulationState.Rand.Int())

	for i := range mappableList {
		immutables := baseSimulation.GenerateRandomImmutables(simulationState.Rand)
		mutables := baseSimulation.GenerateRandomMutables(simulationState.Rand)
		mappableList[i] = mappable.NewClassification(key.NewClassificationID(baseSimulation.GenerateRandomID(simulationState.Rand), immutables, mutables), immutables, mutables)
	}

	genesisState := baseHelpers.NewGenesis(key.Prototype, mappable.Prototype, nil, nil).Initialize(mappableList, []types.Parameter{dummy.Parameter.Mutate(data)})

	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", codec.MustMarshalJSONIndent(simulationState.Cdc, genesisState))

	simulationState.GenState[assetsModule.Name] = simulationState.Cdc.MustMarshalJSON(genesisState)
}
