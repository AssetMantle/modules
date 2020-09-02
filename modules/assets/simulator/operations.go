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

		//	validator, ok := stakingkeeper.RandomValidator(rand, sk, context)
		//	if !ok {
		//		return simulation.NoOpMsg(types.ModuleName), nil, nil // skip
		//	}
		//
		//	simAccount, found := simulation.FindAccount(simulationAccountList, sdkTypes.AccAddress(validator.GetOperator()))
		//	if !found {
		//		return simulation.NoOpMsg(types.ModuleName), nil, nil // skip
		//	}
		//
		//	if !validator.IsJailed() {
		//		return simulation.NoOpMsg(types.ModuleName), nil, nil
		//	}
		//
		//	consAddr := sdkTypes.ConsAddress(validator.GetConsPubKey().Address())
		//	info, found := k.GetValidatorSigningInfo(context, consAddr)
		//	if !found {
		//		return simulation.NoOpMsg(types.ModuleName), nil, nil // skip
		//	}
		//
		//	selfDel := sk.Delegation(context, simAccount.Address, validator.GetOperator())
		//	if selfDel == nil {
		//		return simulation.NoOpMsg(types.ModuleName), nil, nil // skip
		//	}
		//
		//	account := ak.GetAccount(context, sdkTypes.AccAddress(validator.GetOperator()))
		//	fees, err := simulation.RandomFees(rand, context, account.SpendableCoins(context.BlockTime()))
		//	if err != nil {
		//		return simulation.NoOpMsg(types.ModuleName), nil, err
		//	}
		//
		//	msg := types.NewMsgUnjail(validator.GetOperator())
		//
		//	tx := helpers.GenTx(
		//		[]sdkTypes.Msg{msg},
		//		fees,
		//		helpers.DefaultGenTxGas,
		//		chainID,
		//		[]uint64{account.GetAccountNumber()},
		//		[]uint64{account.GetSequence()},
		//		simAccount.PrivKey,
		//	)
		//
		//	_, res, err := baseApp.Deliver(tx)
		//
		//	if info.Tombstoned ||
		//		context.BlockHeader().Time.Before(info.JailedUntil) ||
		//		validator.TokensFromShares(selfDel.GetShares()).TruncateInt().LT(validator.GetMinSelfDelegation()) {
		//		if res != nil && err == nil {
		//			if info.Tombstoned {
		//				return simulation.NewOperationMsg(msg, true, ""), nil, errors.New("validator should not have been unjailed if validator tombstoned")
		//			}
		//			if context.BlockHeader().Time.Before(info.JailedUntil) {
		//				return simulation.NewOperationMsg(msg, true, ""), nil, errors.New("validator unjailed while validator still in jail period")
		//			}
		//			if validator.TokensFromShares(selfDel.GetShares()).TruncateInt().LT(validator.GetMinSelfDelegation()) {
		//				return simulation.NewOperationMsg(msg, true, ""), nil, errors.New("validator unjailed even though self-delegation too low")
		//			}
		//		}
		//		return simulation.NewOperationMsg(msg, false, ""), nil, nil
		//	}
		//
		//	if err != nil {
		//		return simulation.NoOpMsg(types.ModuleName), nil, errors.New(res.Log)
		//	}
		//
		return simulation.NewOperationMsg(nil, true, ""), nil, nil
	}
}
