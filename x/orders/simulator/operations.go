// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypesGo "github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	simulationModules "github.com/AssetMantle/modules/simulation"
	baseTypes "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/simulation/simulated_database/assets"
	"github.com/AssetMantle/modules/simulation/simulated_database/identities"
	"github.com/AssetMantle/modules/simulation/simulated_database/orders"
	"github.com/AssetMantle/modules/x/orders/mappable"
	"github.com/AssetMantle/modules/x/orders/transactions/cancel"
	"github.com/AssetMantle/modules/x/orders/transactions/define"
	"github.com/AssetMantle/modules/x/orders/transactions/make"
	"github.com/AssetMantle/modules/x/orders/transactions/take"
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
			simulateMakeMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateCancelMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateTakeMsg(module),
		),
	}
}

func simulateDefineMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		var identityIDString string

		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		fromIDMap := identities.GetIDData(from.Address.String())

		for _, id := range fromIDMap {
			identityIDString = id
			break
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)
		message := GenerateDefineMessage(from.Address, fromID.(ids.IdentityID), rand)

		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateMakeMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result

		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		message := GetMakeMessage(from, to, rand)
		if message == nil {
			return simulationTypes.NewOperationMsg(&make.Message{}, false, "error in make message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateCancelMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result

		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		makeMessage := GetMakeMessage(from, to, rand)
		if makeMessage == nil {
			return simulationTypes.NewOperationMsg(&make.Message{}, false, "error in make message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		result, err = simulationModules.ExecuteMessage(context, module, makeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(makeMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		orderID := baseIDs.NewOrderID(makeMessage.(*make.Message).ClassificationID, baseQualified.NewImmutables(makeMessage.(*make.Message).ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(makeMessage.(*make.Message).ImmutableProperties.Get()...)...)))

		cancelMessage := cancel.NewMessage(from.Address, makeMessage.(*make.Message).FromID, orderID)
		result, err = simulationModules.ExecuteMessage(context, module, cancelMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(cancelMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(cancelMessage, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateTakeMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result

		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		makeMessage := GetMakeMessage(from, to, rand)
		if makeMessage == nil {
			return simulationTypes.NewOperationMsg(&make.Message{}, false, "error in make message", base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		result, err = simulationModules.ExecuteMessage(context, module, makeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(makeMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		orderID := baseIDs.NewOrderID(makeMessage.(*make.Message).ClassificationID, baseQualified.NewImmutables(makeMessage.(*make.Message).ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(makeMessage.(*make.Message).ImmutableProperties.Get()...)...)))

		takeMessage := take.NewMessage(to.Address, makeMessage.(*make.Message).TakerID, sdkTypes.NewInt(1), orderID)
		result, err = simulationModules.ExecuteMessage(context, module, takeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(takeMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(takeMessage, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}

func GetMakeMessage(from, to simulationTypes.Account, rand *rand.Rand) sdkTypes.Msg {
	var identityIDString, classificationIDString, assetIDString string

	fromIDMap := identities.GetIDData(from.Address.String())
	for _, id := range fromIDMap {
		identityIDString = id
		break
	}
	fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

	toIDMap := identities.GetIDData(to.Address.String())
	for _, id := range toIDMap {
		identityIDString = id
		break
	}

	toID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

	orderMap := orders.GetOrderData(from.Address.String())
	for class, _ := range orderMap {
		classificationIDString = class
		break
	}

	classificationID, _ := baseIDs.PrototypeClassificationID().FromString(classificationIDString)

	assetMap := assets.GetAssetData(from.Address.String())

	for _, id := range assetMap {
		assetIDString = id
	}

	assetID, _ := baseIDs.PrototypeAssetID().FromString(assetIDString)

	mappable := &mappable.Mappable{}
	base.CodecPrototype().Unmarshal(orders.GetMappableBytes(classificationIDString), mappable)
	immutableMetaProperties := &baseLists.PropertyList{}
	immutableProperties := &baseLists.PropertyList{}
	mutableMetaProperties := &baseLists.PropertyList{}
	mutableProperties := &baseLists.PropertyList{}
	if mappable.Order == nil {
		return nil
	}
	for _, i := range mappable.GetOrder().Get().GetImmutables().GetImmutablePropertyList().Get() {
		if i.IsMeta() {
			immutableMetaProperties = immutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
		} else {
			immutableProperties = immutableProperties.Add(i).(*baseLists.PropertyList)
		}
	}
	for _, i := range mappable.GetOrder().Get().GetMutables().GetMutablePropertyList().Get() {
		if i.IsMeta() {
			mutableMetaProperties = mutableMetaProperties.Add(i).(*baseLists.PropertyList)
		} else {
			mutableProperties = mutableProperties.Add(i).(*baseLists.PropertyList)
		}
	}

	return make.NewMessage(from.Address, fromID.(ids.IdentityID), classificationID.(ids.ClassificationID), toID.(ids.IdentityID), assetID.(ids.AssetID), baseDocuments.NewCoinAsset("stake").GetCoinAssetID(), baseTypesGo.NewHeight(-1), sdkTypes.NewInt(1), sdkTypes.NewInt(1), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
}
func GenerateDefineMessage(from sdkTypes.AccAddress, identityID ids.IdentityID, r *rand.Rand) helpers.Message {
	return define.NewMessage(from, identityID, baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r), baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r)).(helpers.Message)
}
