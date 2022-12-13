/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	simulation2 "github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"
)

func (simulator) WeightedProposalContentList() []simulation.WeightedProposalContent {
	return []simulation.WeightedProposalContent{
		simulation2.NewWeightedProposalContent(
			OpWeightSubmitTextProposal,
			DefaultWeightTextProposal,
			contentSimulatorFunc(),
		),
	}
}

func contentSimulatorFunc() simulation.ContentSimulatorFn {
	return func(r *rand.Rand, ctx sdk.Context, accs []simulation.Account) simulation.Content {
		return types.NewTextProposal(
			simulation.RandStringOfLength(r, 140),
			simulation.RandStringOfLength(r, 5000),
		)
	}
}
