/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"math/rand"
)

func WeightedOperations(appParams simulation.AppParams, codec *codec.Codec, transactionKeeper helpers.TransactionKeeper) simulation.WeightedOperations {

	var weightMsg int
	appParams.GetOrGenerate(codec, OpWeightMsg, &weightMsg, nil,
		func(_ *rand.Rand) {
			weightMsg = DefaultWeightMsg
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsg,
			SimulateMsg(transactionKeeper),
		),
	}
}

func SimulateMsg(transactionKeeper helpers.TransactionKeeper) simulation.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulation.Account, chainID string) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		return simulation.NewOperationMsg(nil, true, ""), nil, nil
	}
}
