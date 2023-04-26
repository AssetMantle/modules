// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/AssetMantle/modules/x/classifications/parameters/maxPropertyCount"
	"github.com/AssetMantle/schema/go/properties/utilities"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/classifications/genesis"
	"github.com/AssetMantle/modules/x/classifications/mappable"
	classificationsModule "github.com/AssetMantle/modules/x/classifications/module"
	"github.com/AssetMantle/modules/x/classifications/parameters/bondRate"
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {

	var bondRateData data.Data
	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		bondRate.ID.AsString(),
		&bondRateData,
		simulationState.Rand,
		func(rand *rand.Rand) { bondRateData = baseData.NewNumberData(int64(rand.Intn(99))) },
	)

	var maxPropertyCountData data.Data
	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		maxPropertyCount.ID.AsString(),
		&maxPropertyCountData,
		simulationState.Rand,
		func(rand *rand.Rand) { maxPropertyCountData = baseData.NewNumberData(int64(rand.Intn(99))) },
	)

	mappableList := make([]helpers.Mappable, simulationState.Rand.Intn(99))

	for i := range mappableList {
		immutables := baseQualified.NewImmutables(baseSimulation.GenerateRandomMetaPropertyList(simulationState.Rand).Add(utilities.AnyPropertyListToPropertyList(baseSimulation.GenerateRandomPropertyList(simulationState.Rand).GetList()...)...))
		mutables := baseQualified.NewMutables(baseSimulation.GenerateRandomMetaPropertyList(simulationState.Rand).Add(utilities.AnyPropertyListToPropertyList(baseSimulation.GenerateRandomPropertyList(simulationState.Rand).GetList()...)...))
		mappableList[i] = mappable.NewMappable(base.NewClassification(immutables, mutables))
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseParameters.NewParameterList(bondRate.Parameter.Mutate(bondRateData), maxPropertyCount.Parameter.Mutate(maxPropertyCountData)))

	simulationState.GenState[classificationsModule.Name] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
