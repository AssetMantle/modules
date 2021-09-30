/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gogo/protobuf/proto" //nolint:typecheck
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/common"
	internalGenesis "github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/mappable"
	maintainersModule "github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/parameters/dummy"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/simulation"
	baseSimulation "github.com/persistenceOne/persistenceSDK/simulation/schema/types/base"
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

	mappableList := make([]helpers.Mappable, simulationState.Rand.Intn(99))

	for i := range mappableList {
		mappableList[i] = mappable.NewMaintainer(key.NewMaintainerID(baseSimulation.GenerateRandomID(simulationState.Rand), baseSimulation.GenerateRandomID(simulationState.Rand)), baseSimulation.GenerateRandomProperties(simulationState.Rand), simulation.RandomBool(simulationState.Rand), simulation.RandomBool(simulationState.Rand), simulation.RandomBool(simulationState.Rand))
	}

	genesisState := internalGenesis.NewGenesis(nil, parameters.Prototype().GetList()).Initialize(mappableList, []types.Parameter{dummy.Parameter.Mutate(data)})

	simulationState.GenState[maintainersModule.Name] = common.JSONCodec.MustMarshalJSON(genesisState.(proto.Message))
}
