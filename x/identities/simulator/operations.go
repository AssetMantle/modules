// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	baseTypes "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/simulation/schema/types/database/identities"
	"github.com/AssetMantle/modules/x/identities/transactions/define"
	"github.com/AssetMantle/modules/x/identities/transactions/issue"
	"github.com/AssetMantle/modules/x/identities/transactions/nub"
	"github.com/AssetMantle/schema/go/ids"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) WeightedOperations(simulationState module.SimulationState, module helpers.Module) simulation.WeightedOperations {
	var weightMsg int

	simulationState.AppParams.GetOrGenerate(nil, OpWeightMsg, &weightMsg, nil,
		func(_ *rand.Rand) {
			weightMsg = DefaultWeightMsg
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsg+1000,
			simulateNubMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg+100,
			simulateDefineMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateIssueMsg(module),
		),
	}
}

func ExecuteMessage(context sdkTypes.Context, module helpers.Module, message helpers.Message) (*sdkTypes.Result, error) {
	return module.GetTransactions().Get(message.Type()).HandleMessage(sdkTypes.WrapSDKContext(context), message)
}

func simulateNubMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		account, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		message := GenerateNubMessage(account.Address, baseTypes.GenerateRandomID(rand))
		result, err := ExecuteMessage(context, module, message)
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, "error executing Nub message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		identities.AddAccountNubIDPair(account.Address, string(result.Data))
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateDefineMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		var message helpers.Message
		account, nubID := identities.GetRandomAccNubIDPair()
		message = GenerateDefineMessage(account, nubID, rand)
		result, err = ExecuteMessage(context, module, message)
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, "error executing define message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		identities.AddAccountClassificationIDPair(account, string(result.Data))
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateIssueMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		var message helpers.Message
		from, nubID := identities.GetRandomAccNubIDPair()
		//to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		message = GenerateIssueMessage(from, from, nubID, identities.GetClassificationID(from), rand)
		result, err = ExecuteMessage(context, module, message)
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, "error executing define message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		identities.AddIssuedIdentity(from, string(result.Data), from)
		x := identities.GetIssuedIdentityInfo(from)
		fmt.Println(x)
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}

func GenerateNubMessage(from sdkTypes.AccAddress, nubID ids.ID) helpers.Message {
	return nub.NewMessage(from, nubID).(helpers.Message)
}
func GenerateDefineMessage(from sdkTypes.AccAddress, nubID ids.IdentityID, r *rand.Rand) helpers.Message {
	return define.NewMessage(from, nubID, baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r), baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r)).(helpers.Message)
}
func GenerateIssueMessage(from sdkTypes.AccAddress, to sdkTypes.AccAddress, fromID ids.IdentityID, classificationID ids.ClassificationID, r *rand.Rand) helpers.Message {
	return issue.NewMessage(from, to, fromID, classificationID, baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r), baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r)).(helpers.Message)
}
