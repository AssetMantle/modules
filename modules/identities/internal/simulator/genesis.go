// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/modules/identities/internal/common"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	identitiesModule "github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters/dummy"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	parameters2 "github.com/AssetMantle/modules/schema/parameters"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		dummy.ID.String(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = base.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, len(simulationState.Accounts))

	for i := range mappableList {
		immutableProperties := baseSimulation.GenerateRandomProperties(simulationState.Rand)
		mappableList[i] = mappable.NewIdentity(key.NewIdentityID(baseSimulation.GenerateRandomID(simulationState.Rand), immutableProperties), immutableProperties, baseSimulation.GenerateRandomProperties(simulationState.Rand))
	}

	genesisState := baseHelpers.NewGenesis(key.Prototype, mappable.Prototype, nil, parameters.Prototype().GetList()).Initialize(mappableList, []parameters2.Parameter{dummy.Parameter.Mutate(Data)})

	simulationState.GenState[identitiesModule.Name] = common.Codec.MustMarshalJSON(genesisState)
}
