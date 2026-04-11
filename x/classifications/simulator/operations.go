// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
)

// WeightedOperations returns nil because the classifications module only exposes
// a govern transaction (parameter changes via governance proposals). Classification
// logic is exercised transitively when other modules' operations trigger
// classifications auxiliaries (define, bond, conform, etc.).
func (simulator) WeightedOperations(simulationState module.SimulationState, module helpers.Module) simulation.WeightedOperations {
	var weightMsg int

	simulationState.AppParams.GetOrGenerate(OpWeightMsg, &weightMsg, nil,
		func(_ *rand.Rand) {
			weightMsg = DefaultWeightMsg
		},
	)

	return nil
}
