// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) WeightedOperations(appParams simulationTypes.AppParams, jsonCodec codec.JSONCodec) simulation.WeightedOperations {
	var weightMsg int

	appParams.GetOrGenerate(jsonCodec, OpWeightMsg, &weightMsg, nil,
		func(_ *rand.Rand) {
			weightMsg = DefaultWeightMsg
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsg,
			simulateMsg(),
		),
	}
}

func simulateMsg() simulation.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulation.Account, chainID string) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		return simulation.NewOperationMsg(nil, true, ""), nil, nil
	}
}
