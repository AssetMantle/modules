/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) WeightedOperations(appParams simulation.AppParams, codec *codec.Codec) simulation.WeightedOperations {
	var weightMsg int

	appParams.GetOrGenerate(codec, OpWeightMsg, &weightMsg, nil,
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
