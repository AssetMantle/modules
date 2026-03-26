// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"
)

func (simulator) ProposalMessages(_ module.SimulationState) []simulationTypes.WeightedProposalMsg {
	return []simulationTypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			OpWeightSubmitTextProposal,
			DefaultWeightTextProposal,
			func(r *rand.Rand, _ sdkTypes.Context, simulationAccounts []simulationTypes.Account) sdkTypes.Msg {
				msgSubmitProposal, err := v1beta1.NewMsgSubmitProposal(v1beta1.NewTextProposal(simulationTypes.RandStringOfLength(r, 140), simulationTypes.RandStringOfLength(r, 5000)), sdkTypes.NewCoins(sdkTypes.NewInt64Coin(sdkTypes.DefaultBondDenom, 10000000)), simulationAccounts[r.Intn(len(simulationAccounts))].Address)
				if err != nil {
					panic(err)
				}

				return msgSubmitProposal
			},
		),
	}
}
