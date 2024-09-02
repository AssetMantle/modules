// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"
)

type Simulator interface {
	RandomizedGenesisState(*module.SimulationState)
	WeightedOperations(module.SimulationState, Module) simulation.WeightedOperations
	ParamChangeList(*rand.Rand) []simulationTypes.LegacyParamChange
	ProposalMessages(module.SimulationState) []simulationTypes.WeightedProposalMsg
}
