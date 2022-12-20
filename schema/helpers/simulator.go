// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

type Simulator interface {
	RandomizedGenesisState(*module.SimulationState)
	WeightedOperations(module.SimulationState) simulation.WeightedOperations
	WeightedProposalContentList(module.SimulationState) []simulationTypes.WeightedProposalContent
	ParamChangeList(*rand.Rand) []simulationTypes.ParamChange
}
