// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	base2 "github.com/AssetMantle/modules/helpers/base"
	"math/rand"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/AssetMantle/modules/helpers"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/splits/internal/genesis"
	"github.com/AssetMantle/modules/x/splits/internal/mappable"
	splitsModule "github.com/AssetMantle/modules/x/splits/internal/module"
	"github.com/AssetMantle/modules/x/splits/internal/parameters/wrapAllowedCoins"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		wrapAllowedCoins.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, simulationState.Rand.Intn(99))

	for i := range mappableList {
		immutables := baseQualified.NewImmutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand))
		mutables := baseQualified.NewMutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand))
		mappableList[i] = mappable.NewMappable(base.NewSplit(baseIDs.NewIdentityID(baseIDs.NewClassificationID(immutables, mutables), immutables), baseIDs.NewCoinID(baseIDs.NewStringID(simulationTypes.RandStringOfLength(simulationState.Rand, simulationState.Rand.Intn(99)))), simulationTypes.RandomDecAmount(simulationState.Rand, sdkTypes.NewDec(9999999999))))
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseParameters.NewParameterList(wrapAllowedCoins.Parameter.Mutate(Data)))

	simulationState.GenState[splitsModule.Name] = base2.CodecPrototype().MustMarshalJSON(genesisState)
}
