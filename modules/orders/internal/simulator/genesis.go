/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gogo/protobuf/proto" //nolint:typecheck
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/common"
	internalGenesis "github.com/persistenceOne/persistenceSDK/modules/orders/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mappable"
	ordersModule "github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/parameters/dummy"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
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
		immutables := baseSimulation.GenerateRandomProperties(simulationState.Rand)
		mappableList[i] = mappable.NewOrder(key.NewOrderID(baseSimulation.GenerateRandomID(simulationState.Rand), baseSimulation.GenerateRandomID(simulationState.Rand), baseSimulation.GenerateRandomID(simulationState.Rand), baseSimulation.GenerateRandomIDWithDec(simulationState.Rand), baseSimulation.GenerateRandomIDWithInt64(simulationState.Rand), baseSimulation.GenerateRandomID(simulationState.Rand), immutables), immutables, baseSimulation.GenerateRandomProperties(simulationState.Rand))
	}
	parametersList := parameters.Prototype().GetList()
	newParametersList := make([]dummy.DummyParameter, len(parametersList))
	for i, _ := range parametersList {
		newParametersList[i] = *dummy.NewParameter(parametersList[i].GetID(), parametersList[i].GetData())
	}
	genesisState := internalGenesis.NewGenesis(nil, newParametersList).Initialize(mappableList, []types.Parameter{dummy.Parameter.Mutate(data)})

	simulationState.GenState[ordersModule.Name] = common.JSONCodec.MustMarshalJSON(genesisState.(proto.Message))
}
