// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/AssetMantle/modules/helpers/base"
	simulationModules "github.com/AssetMantle/modules/simulation"
	baseTypes "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/identities"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/transactions/define"
	"github.com/AssetMantle/modules/x/identities/transactions/issue"
	"github.com/AssetMantle/modules/x/identities/transactions/nub"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
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
			weightMsg+1000,
			simulateDefineMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg+1000,
			simulateIssueMsg(module),
		),
	}
}

func simulateNubMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		account, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		message := GenerateNubMessage(account.Address, baseTypes.GenerateRandomID(rand))
		result, err := simulationModules.ExecuteMessage(context, module, message)
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, "error executing Nub message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateDefineMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		var message *define.Message
		var identityIDString string
		account, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		identityMap := identities.GetIDData(account.Address.String())
		for _, id := range identityMap {
			identityIDString = id
			break
		}
		identityID, _ := baseIDs.ReadIdentityID(identityIDString)
		message = GenerateDefineMessage(account.Address, identityID, rand).(*define.Message)
		result, err = simulationModules.ExecuteMessage(context, module, message)
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, "error executing define message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}

func simulateIssueMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		var classificationIDString, identityIDString string
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		identityMap := identities.GetIDData(from.Address.String())
		for class, id := range identityMap {
			identityIDString = id
			classificationIDString = class
			break
		}
		identityID, _ := baseIDs.ReadIdentityID(identityIDString)
		classificationID, _ := baseIDs.ReadClassificationID(classificationIDString)
		mappable := &mappable.Mappable{}
		base.CodecPrototype().Unmarshal(identities.GetMappableBytes(classificationIDString), mappable)
		immutableMetaProperties := &baseLists.PropertyList{}
		immutableProperties := &baseLists.PropertyList{}
		mutableMetaProperties := &baseLists.PropertyList{}
		mutableProperties := &baseLists.PropertyList{}
		if mappable.Identity == nil {
			return simulationTypes.NewOperationMsg(&issue.Message{}, false, "invalid identity", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		for _, i := range mappable.GetIdentity().Get().GetImmutables().GetImmutablePropertyList().GetList() {
			if i.IsMeta() {
				immutableMetaProperties = immutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
			} else {
				immutableProperties = immutableProperties.Add(i).(*baseLists.PropertyList)
			}
		}
		for _, i := range mappable.GetIdentity().Get().GetMutables().GetMutablePropertyList().GetList() {
			if i.IsMeta() {
				mutableMetaProperties = mutableMetaProperties.Add(i).(*baseLists.PropertyList)
			} else {
				mutableProperties = mutableProperties.Add(i).(*baseLists.PropertyList)
			}
		}

		message := issue.NewMessage(from.Address, to.Address, identityID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)

		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, "error executing issue message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
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
