// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) WeightedProposalContentList(module.SimulationState) []simulationTypes.WeightedProposalContent {
	return []simulationTypes.WeightedProposalContent{
		simulation.NewWeightedProposalContent(
			OpWeightSubmitTextProposal,
			DefaultWeightTextProposal,
			simulateTextProposalContent,
		),
	}
}

func simulateTextProposalContent(r *rand.Rand, _ sdkTypes.Context, _ []simulationTypes.Account) simulationTypes.Content {
	return types.NewTextProposal(
		simulationTypes.RandStringOfLength(r, 140),
		simulationTypes.RandStringOfLength(r, 5000),
	)
}
