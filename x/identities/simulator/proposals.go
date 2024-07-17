// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	govSimulation "github.com/cosmos/cosmos-sdk/x/gov/simulation"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) ProposalMessages(_ module.SimulationState) []simulationTypes.WeightedProposalMsg {
	return []simulationTypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			OpWeightSubmitTextProposal,
			DefaultWeightTextProposal,
			func(r *rand.Rand, _ sdkTypes.Context, simulationAccounts []simulationTypes.Account) sdkTypes.Msg {
				msgSubmitProposal, err := v1beta1.NewMsgSubmitProposal(v1beta1.NewTextProposal(simulationTypes.RandStringOfLength(r, 140), simulationTypes.RandStringOfLength(r, 5000)), govSimulation.GenDepositParamsMinDeposit(r), simulationAccounts[r.Intn(len(simulationAccounts))].Address)
				if err != nil {
					panic(err)
				}

				return msgSubmitProposal
			},
		),
	}
}
