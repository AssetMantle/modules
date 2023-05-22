// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/classifications/genesis"
	"github.com/AssetMantle/modules/x/classifications/mappable"
	classificationsModule "github.com/AssetMantle/modules/x/classifications/module"
	"github.com/AssetMantle/modules/x/classifications/parameters/bondRate"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		bondRate.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, simulationState.Rand.Intn(99))

	for i := range mappableList {
		immutables := baseQualified.NewImmutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand))
		mutables := baseQualified.NewMutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand))
		mappableList[i] = mappable.NewMappable(base.NewClassification(immutables, mutables))
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseLists.NewParameterList(bondRate.Parameter.Mutate(Data)))

	simulationState.GenState[classificationsModule.Name] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
