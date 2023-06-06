// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math"
	"math/rand"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/lists/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/metas/constants"
	"github.com/AssetMantle/modules/x/metas/genesis"
	"github.com/AssetMantle/modules/x/metas/mappable"
	"github.com/AssetMantle/modules/x/metas/parameters/revealEnabled"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		revealEnabled.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, simulationState.Rand.Intn(99))

	for i := range mappableList {
		mappableList[i] = mappable.NewMappable(baseSimulation.GenerateRandomData(simulationState.Rand, int(math.Abs(float64(simulationState.Rand.Int())))))
	}

	genesisState := genesis.Prototype().Initialize(mappableList, base.NewParameterList(revealEnabled.Parameter.Mutate(Data)))

	simulationState.GenState[constants.ModuleName] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
