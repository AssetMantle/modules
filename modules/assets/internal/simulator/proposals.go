/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) WeightedProposalContentList() []simulation.WeightedProposalContent {
	return []simulation.WeightedProposalContent{
		{
			AppParamsKey:       OpWeightSubmitTextProposal,
			DefaultWeight:      DefaultWeightTextProposal,
			ContentSimulatorFn: simulateTextProposalContent,
		},
	}
}

func simulateTextProposalContent(r *rand.Rand, _ sdk.Context, _ []simulation.Account) types.Content {
	return types.NewTextProposal(
		simulation.RandStringOfLength(r, 140),
		simulation.RandStringOfLength(r, 5000),
	)
}
