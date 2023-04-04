// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/x/data"
	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/documents/base"
	"github.com/AssetMantle/schema/x/helpers"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseParameters "github.com/AssetMantle/schema/x/parameters/base"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/modules/assets/internal/common"
	"github.com/AssetMantle/modules/modules/assets/internal/genesis"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	assetsModule "github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters/mintEnabled"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		mintEnabled.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, simulationState.Rand.Intn(99))

	for i := range mappableList {
		immutables := baseQualified.NewImmutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand))
		mutables := baseQualified.NewMutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand))
		mappableList[i] = mappable.NewMappable(base.NewAsset(baseIDs.NewClassificationID(immutables, mutables), immutables, mutables))
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseParameters.NewParameterList(mintEnabled.Parameter.Mutate(Data)))

	simulationState.GenState[assetsModule.Name] = common.LegacyAmino.MustMarshalJSON(genesisState)
}
