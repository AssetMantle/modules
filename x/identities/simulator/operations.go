// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/AssetMantle/modules/helpers/base"
	simulationModules "github.com/AssetMantle/modules/simulation"
	baseTypes "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/identities"
	"github.com/AssetMantle/modules/x/assets/transactions/revoke"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/transactions/define"
	"github.com/AssetMantle/modules/x/identities/transactions/deputize"
	"github.com/AssetMantle/modules/x/identities/transactions/issue"
	"github.com/AssetMantle/modules/x/identities/transactions/mutate"
	"github.com/AssetMantle/modules/x/identities/transactions/nub"
	"github.com/AssetMantle/modules/x/identities/transactions/provision"
	"github.com/AssetMantle/modules/x/identities/transactions/quash"
	"github.com/AssetMantle/modules/x/identities/transactions/unprovision"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/utilities"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
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
			weightMsg,
			simulateNubMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateDefineMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateIssueMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateProvisionAndUnprovisionMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateDeputizeAndRevokeMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateQuashMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateMutateMsg(module),
		),
	}
}

func simulateNubMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		account, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		message := GenerateNubMessage(account.Address, baseTypes.GenerateRandomID(rand))
		result, err := simulationModules.ExecuteMessage(context, module, message)
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
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
			return simulationTypes.NewOperationMsg(message, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateIssueMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		message := GetIssueMessage(from, to, rand)
		if message == nil {
			return simulationTypes.NewOperationMsg(&issue.Message{}, false, "error in issue message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateProvisionAndUnprovisionMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		var identityIDString string
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		identityMap := identities.GetIDData(from.Address.String())
		for _, id := range identityMap {
			identityIDString = id
			break
		}
		identityID, _ := baseIDs.ReadIdentityID(identityIDString)

		provisionMessage := provision.NewMessage(from.Address, to.Address, identityID)
		result, err = simulationModules.ExecuteMessage(context, module, provisionMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(provisionMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		unprovisionMessage := unprovision.NewMessage(from.Address, to.Address, identityID)
		result, err = simulationModules.ExecuteMessage(context, module, unprovisionMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(unprovisionMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(unprovisionMessage, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateDeputizeAndRevokeMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		var classificationIDString, identityIDString string
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		fromIDMap := identities.GetIDData(from.Address.String())

		if fromIDMap == nil {
			return simulationTypes.NewOperationMsg(&deputize.Message{}, false, "address not found", base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		for class, id := range fromIDMap {
			identityIDString = id
			classificationIDString = class
			break

		}

		classificationID, _ := baseIDs.ReadClassificationID(classificationIDString)
		fromID, _ := baseIDs.ReadIdentityID(identityIDString)
		toIDMap := identities.GetIDData(to.Address.String())

		if toIDMap == nil {
			return simulationTypes.NewOperationMsg(&deputize.Message{}, false, "address not found", base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		for _, id := range toIDMap {
			identityIDString = id
			break
		}

		toID, _ := baseIDs.ReadIdentityID(identityIDString)
		mappable := &mappable.Mappable{}
		base.CodecPrototype().Unmarshal(identities.GetMappableBytes(classificationIDString), mappable)
		deputizeMessage := deputize.NewMessage(from.Address, fromID, toID, classificationID, mappable.Identity.Mutables.PropertyList, true, true, true, true, true, true)
		result, err = simulationModules.ExecuteMessage(context, module, deputizeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(deputizeMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		revokeMessage := revoke.NewMessage(from.Address, fromID, toID, classificationID)
		result, err = simulationModules.ExecuteMessage(context, module, revokeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(revokeMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(revokeMessage, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateQuashMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		message := GetIssueMessage(from, to, rand)
		if message == nil {
			return simulationTypes.NewOperationMsg(&issue.Message{}, false, "error in issue message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		id := baseIDs.NewIdentityID(message.(*issue.Message).ClassificationID, baseQualified.NewImmutables(message.(*issue.Message).ImmutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.(*issue.Message).ImmutableProperties.GetList()...)...)))
		quashMessage := quash.NewMessage(from.Address, message.(*issue.Message).FromID, id)

		result, err = simulationModules.ExecuteMessage(context, module, quashMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(quashMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(quashMessage, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateMutateMsg(module helpers.Module) simulationTypes.Operation {
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
		fromID, _ := baseIDs.ReadIdentityID(identityIDString)
		classificationID, _ := baseIDs.ReadClassificationID(classificationIDString)
		mappable := &mappable.Mappable{}
		base.CodecPrototype().Unmarshal(identities.GetMappableBytes(classificationIDString), mappable)
		immutableMetaProperties := &baseLists.PropertyList{}
		immutableProperties := &baseLists.PropertyList{}
		mutableMetaProperties := &baseLists.PropertyList{}
		mutableProperties := &baseLists.PropertyList{}
		updatedProperties := &baseLists.PropertyList{}
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
				updatedProperties = mutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
			} else {
				mutableProperties = mutableProperties.Add(i).(*baseLists.PropertyList)
			}
		}

		message := issue.NewMessage(from.Address, to.Address, fromID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)

		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		issuedID := baseIDs.NewIdentityID(classificationID, baseQualified.NewImmutables(immutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(immutableProperties.GetList()...)...)))
		mutateMessage := mutate.NewMessage(from.Address, fromID, issuedID, updatedProperties, mutableProperties)

		result, err = simulationModules.ExecuteMessage(context, module, mutateMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(mutateMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(mutateMessage, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}

func GenerateNubMessage(from sdkTypes.AccAddress, nubID ids.ID) helpers.Message {
	return nub.NewMessage(from, nubID).(helpers.Message)
}
func GenerateDefineMessage(from sdkTypes.AccAddress, nubID ids.IdentityID, r *rand.Rand) helpers.Message {
	return define.NewMessage(from, nubID, baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r), baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r)).(helpers.Message)
}
func GetIssueMessage(from, to simulationTypes.Account, rand *rand.Rand) sdkTypes.Msg {
	var classificationIDString, identityIDString string

	identityMap := identities.GetIDData(from.Address.String())
	for class, id := range identityMap {
		identityIDString = id
		classificationIDString = class
		break
	}
	fromID, _ := baseIDs.ReadIdentityID(identityIDString)
	classificationID, _ := baseIDs.ReadClassificationID(classificationIDString)
	mappable := &mappable.Mappable{}
	base.CodecPrototype().Unmarshal(identities.GetMappableBytes(classificationIDString), mappable)
	immutableMetaProperties := &baseLists.PropertyList{}
	immutableProperties := &baseLists.PropertyList{}
	mutableMetaProperties := &baseLists.PropertyList{}
	mutableProperties := &baseLists.PropertyList{}
	if mappable.Identity == nil {
		return nil
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
	return issue.NewMessage(from.Address, to.Address, fromID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
}
