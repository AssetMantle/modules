// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/modules/modules/maintainers/internal/utilities"
	"github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/utilities/random"

	"github.com/AssetMantle/modules/modules/maintainers/internal/common"
	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/internal/mappable"
	maintainersModule "github.com/AssetMantle/modules/modules/maintainers/internal/module"
	"github.com/AssetMantle/modules/modules/maintainers/internal/parameters"
	"github.com/AssetMantle/modules/modules/maintainers/internal/parameters/dummy"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	parametersSchema "github.com/AssetMantle/modules/schema/parameters"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		dummy.ID.String(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, simulationState.Rand.Intn(99))

	var classificationID ids.ClassificationID
	for i := range mappableList {
		immutables := baseQualified.NewImmutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand))
		mutables := baseQualified.NewMutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand))
		classificationID = baseIDs.NewClassificationID(immutables, mutables)
		mappableList[i] = mappable.NewMappable(base.NewMaintainer(baseIDs.NewIdentityID(classificationID, immutables), classificationID, mutables.GetMutablePropertyList().GetPropertyIDList(), utilities.SetPermissions(random.GenerateRandomBool(), random.GenerateRandomBool(), random.GenerateRandomBool(), random.GenerateRandomBool(), random.GenerateRandomBool(), random.GenerateRandomBool())))
	}

	genesisState := baseHelpers.NewGenesis(key.Prototype, mappable.Prototype, nil, parameters.Prototype().GetList()).Initialize(mappableList, []parametersSchema.Parameter{dummy.Parameter.Mutate(Data)})

	simulationState.GenState[maintainersModule.Name] = common.Codec.MustMarshalJSON(genesisState)
}
