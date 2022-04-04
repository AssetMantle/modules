// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/modules/orders/internal/common"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	ordersModule "github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/orders/internal/parameters"
	"github.com/AssetMantle/modules/modules/orders/internal/parameters/dummy"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
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

	mappableList := make([]helpers.Mappable, simulationState.Rand.Intn(99))

	for i := range mappableList {
		immutables := baseSimulation.GenerateRandomProperties(simulationState.Rand)
		mappableList[i] = mappable.NewOrder(key.NewOrderID(baseSimulation.GenerateRandomID(simulationState.Rand), baseSimulation.GenerateRandomID(simulationState.Rand), baseSimulation.GenerateRandomID(simulationState.Rand), baseSimulation.GenerateRandomIDWithDec(simulationState.Rand), baseSimulation.GenerateRandomIDWithInt64(simulationState.Rand), baseSimulation.GenerateRandomID(simulationState.Rand), immutables), immutables, baseSimulation.GenerateRandomProperties(simulationState.Rand))
	}

	genesisState := baseHelpers.NewGenesis(key.Prototype, mappable.Prototype, nil, parameters.Prototype().GetList()).Initialize(mappableList, []types.Parameter{dummy.Parameter.Mutate(data)})

	simulationState.GenState[ordersModule.Name] = common.Codec.MustMarshalJSON(genesisState)
}
