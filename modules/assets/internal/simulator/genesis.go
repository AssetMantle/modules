/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"fmt"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"math"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mappable"
	assetsModule "github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters/dummy"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
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

	var mappableList []helpers.Mappable
	for range simulationState.Accounts {
		immutables := GenerateRandomImmutables(simulationState.Rand)
		mappableList = append(mappableList, mappable.NewAsset(key.NewAssetID(GenerateRandomID(simulationState.Rand), immutables), immutables, GenerateRandomMutables(simulationState.Rand)))
	}
	genesisState := baseHelpers.NewGenesis(key.Prototype, mappable.Prototype, nil, nil).Initialize(mappableList, []types.Parameter{dummy.Parameter.Mutate(data)})

	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", codec.MustMarshalJSONIndent(simulationState.Cdc, genesisState))
	simulationState.GenState[assetsModule.Name] = simulationState.Cdc.MustMarshalJSON(genesisState)
}

func GenerateRandomImmutables(r *rand.Rand) types.Immutables {
	return base.NewImmutables(GenerateRandomProperties(r))
}

func GenerateRandomMutables(r *rand.Rand) types.Mutables {
	return base.NewMutables(GenerateRandomProperties(r))
}

func GenerateRandomID(r *rand.Rand) types.ID {
	return base.NewID(simulation.RandStringOfLength(r, r.Int()))
}

func GenerateRandomFact(r *rand.Rand) types.Fact {
	randomPositiveInt := int(math.Abs(float64(r.Int())))
	var data types.Data
	switch randomPositiveInt % 4 {
	case 0:
		data = base.NewIDData(GenerateRandomID(r))
	case 1:
		data = base.NewStringData(simulation.RandStringOfLength(r, r.Int()))
	case 2:
		data = base.NewDecData(simulation.RandomDecAmount(r, sdkTypes.NewDec(9999999999)))
	case 3:
		data = base.NewHeightData(base.NewHeight(r.Int63()))
	default:
		return nil
	}
	return base.NewFact(data)
}

func GenerateRandomProperty(r *rand.Rand) types.Property {
	return base.NewProperty(GenerateRandomID(r), GenerateRandomFact(r))
}

func GenerateRandomProperties(r *rand.Rand) types.Properties {
	randomPositiveInt := int(math.Abs(float64(r.Int()))) % 11
	var propertyList []types.Property
	for i := 0; i <= randomPositiveInt; i++ {
		propertyList = append(propertyList, GenerateRandomProperty(r))
	}
	return base.NewProperties(propertyList...)
}
