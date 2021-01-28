/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mappable"
	assetsModule "github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters/dummy"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	baseSimulation "github.com/persistenceOne/persistenceSDK/utilities/simulation/schema/types/base"
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

	mappableList := make([]helpers.Mappable, len(simulationState.Accounts))

	for i := range simulationState.Accounts {
		immutables := baseSimulation.GenerateRandomImmutables(simulationState.Rand)
		mappableList[i] = mappable.NewAsset(key.NewAssetID(baseSimulation.GenerateRandomID(simulationState.Rand), immutables), immutables, baseSimulation.GenerateRandomMutables(simulationState.Rand))
	}
	genesisState := baseHelpers.NewGenesis(key.Prototype, mappable.Prototype, nil, nil).Initialize(mappableList, []types.Parameter{dummy.Parameter.Mutate(data)})

	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", codec.MustMarshalJSONIndent(simulationState.Cdc, genesisState))
	simulationState.GenState[assetsModule.Name] = simulationState.Cdc.MustMarshalJSON(genesisState)
}
