// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/AssetMantle/modules/helpers/base"
	simulationModules "github.com/AssetMantle/modules/simulation"
	baseTypes "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/simulation/simulated_database/assets"
	"github.com/AssetMantle/modules/simulation/simulated_database/identities"
	"github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/assets/transactions/burn"
	"github.com/AssetMantle/modules/x/assets/transactions/define"
	"github.com/AssetMantle/modules/x/assets/transactions/deputize"
	"github.com/AssetMantle/modules/x/assets/transactions/mint"
	"github.com/AssetMantle/modules/x/assets/transactions/mutate"
	"github.com/AssetMantle/modules/x/assets/transactions/renumerate"
	"github.com/AssetMantle/modules/x/assets/transactions/revoke"
	"github.com/AssetMantle/modules/x/identities/transactions/issue"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/AssetMantle/modules/helpers"
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
			simulateDefineMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateMintMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateBurnMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateRenumerateMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateDeputizeAndRevokeMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateMutateMsg(module),
		),
	}
	return nil
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
func simulateMintMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		message := GetMintMessage(from, to, rand)
		if message == nil {
			return simulationTypes.NewOperationMsg(&mint.Message{}, false, "error in mint message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateRenumerateMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		mintMessage := GetMintMessage(from, to, rand)
		if mintMessage == nil {
			return simulationTypes.NewOperationMsg(&mint.Message{}, false, "error in mint message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		result, err = simulationModules.ExecuteMessage(context, module, mintMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(&renumerate.Message{}, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		assetID := baseIDs.NewAssetID(mintMessage.(*mint.Message).ClassificationID, baseQualified.NewImmutables(mintMessage.(*mint.Message).ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(mintMessage.(*mint.Message).ImmutableProperties.Get()...)...)))
		renumerateMessage := renumerate.NewMessage(from.Address, mintMessage.(*mint.Message).FromID, assetID)
		result, err = simulationModules.ExecuteMessage(context, module, renumerateMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(renumerateMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(renumerateMessage, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateBurnMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		mintMessage := GetMintMessage(from, to, rand)
		if mintMessage == nil {
			return simulationTypes.NewOperationMsg(&mint.Message{}, false, "error in mint message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		result, err = simulationModules.ExecuteMessage(context, module, mintMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(mintMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		assetID := baseIDs.NewAssetID(mintMessage.(*mint.Message).ClassificationID, baseQualified.NewImmutables(mintMessage.(*mint.Message).ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(mintMessage.(*mint.Message).ImmutableProperties.Get()...)...)))
		burnMessage := burn.NewMessage(from.Address, mintMessage.(*mint.Message).FromID, assetID)
		result, err = simulationModules.ExecuteMessage(context, module, burnMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(burnMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(burnMessage, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
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

		for _, id := range fromIDMap {
			identityIDString = id
			break

		}

		assetMap := assets.GetAssetData(from.Address.String())

		for class, _ := range assetMap {
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
		base.CodecPrototype().Unmarshal(assets.GetMappableBytes(classificationIDString), mappable)
		deputizeMessage := deputize.NewMessage(from.Address, fromID, toID, classificationID, mappable.Asset.Mutables.PropertyList, true, true, true, true, true, true)
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
func simulateMutateMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		var classificationIDString, identityIDString string
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		fromIDMap := identities.GetIDData(from.Address.String())
		for _, id := range fromIDMap {
			identityIDString = id
			break
		}
		fromID, _ := baseIDs.ReadIdentityID(identityIDString)

		toIDMap := identities.GetIDData(to.Address.String())
		for _, id := range toIDMap {
			identityIDString = id
			break
		}
		toID, _ := baseIDs.ReadIdentityID(identityIDString)

		assetMap := assets.GetAssetData(from.Address.String())
		for class, _ := range assetMap {
			classificationIDString = class
			break
		}
		classificationID, _ := baseIDs.ReadClassificationID(classificationIDString)

		mappable := &mappable.Mappable{}
		base.CodecPrototype().Unmarshal(assets.GetMappableBytes(classificationIDString), mappable)
		immutableMetaProperties := &baseLists.PropertyList{}
		immutableProperties := &baseLists.PropertyList{}
		mutableMetaProperties := &baseLists.PropertyList{}
		mutableProperties := &baseLists.PropertyList{}
		updatedProperties := &baseLists.PropertyList{}
		if mappable.Asset == nil {
			return simulationTypes.NewOperationMsg(&issue.Message{}, false, "invalid identity", base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		for _, i := range mappable.GetAsset().Get().GetImmutables().GetImmutablePropertyList().Get() {
			if i.IsMeta() {
				immutableMetaProperties = immutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
			} else {
				immutableProperties = immutableProperties.Add(i).(*baseLists.PropertyList)
			}
		}
		for _, i := range mappable.GetAsset().Get().GetMutables().GetMutablePropertyList().Get() {
			if i.IsMeta() {
				mutableMetaProperties = mutableMetaProperties.Add(i).(*baseLists.PropertyList)
				updatedProperties = mutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
			} else {
				mutableProperties = mutableProperties.Add(i).(*baseLists.PropertyList)
			}
		}

		mintMessage := mint.NewMessage(from.Address, fromID, toID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
		result, err = simulationModules.ExecuteMessage(context, module, mintMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(mintMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		mintedAssetID := baseIDs.NewAssetID(classificationID, baseQualified.NewImmutables(immutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(immutableProperties.Get()...)...)))
		mutateMessage := mutate.NewMessage(from.Address, fromID, mintedAssetID, updatedProperties, mutableProperties)

		result, err = simulationModules.ExecuteMessage(context, module, mutateMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(mutateMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(mutateMessage, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}

func GenerateDefineMessage(from sdkTypes.AccAddress, identityID ids.IdentityID, r *rand.Rand) helpers.Message {
	return define.NewMessage(from, identityID, baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r), baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r)).(helpers.Message)
}
func GetMintMessage(from, to simulationTypes.Account, rand *rand.Rand) sdkTypes.Msg {
	var classificationIDString, identityIDString string

	fromIDMap := identities.GetIDData(from.Address.String())
	for _, id := range fromIDMap {
		identityIDString = id
		break
	}
	fromID, _ := baseIDs.ReadIdentityID(identityIDString)

	toIDMap := identities.GetIDData(to.Address.String())
	for _, id := range toIDMap {
		identityIDString = id
		break
	}
	toID, _ := baseIDs.ReadIdentityID(identityIDString)

	assetMap := assets.GetAssetData(from.Address.String())
	for class, _ := range assetMap {
		classificationIDString = class
		break
	}
	classificationID, _ := baseIDs.ReadClassificationID(classificationIDString)
	mappable := &mappable.Mappable{}
	base.CodecPrototype().Unmarshal(assets.GetMappableBytes(classificationIDString), mappable)
	immutableMetaProperties := &baseLists.PropertyList{}
	immutableProperties := &baseLists.PropertyList{}
	mutableMetaProperties := &baseLists.PropertyList{}
	mutableProperties := &baseLists.PropertyList{}
	if mappable.Asset == nil {
		return nil
	}
	for _, i := range mappable.GetAsset().Get().GetImmutables().GetImmutablePropertyList().Get() {
		if i.IsMeta() {
			immutableMetaProperties = immutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
		} else {
			immutableProperties = immutableProperties.Add(i).(*baseLists.PropertyList)
		}
	}
	for _, i := range mappable.GetAsset().Get().GetMutables().GetMutablePropertyList().Get() {
		if i.IsMeta() {
			mutableMetaProperties = mutableMetaProperties.Add(i).(*baseLists.PropertyList)
		} else {
			mutableProperties = mutableProperties.Add(i).(*baseLists.PropertyList)
		}
	}
	return mint.NewMessage(from.Address, fromID, toID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
}
