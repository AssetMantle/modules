// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"cosmossdk.io/math"
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
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	simulationModules "github.com/AssetMantle/modules/simulation"
	baseTypes "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/simulation/simulated_database/orders"
	"github.com/AssetMantle/modules/x/orders/mappable"
	"github.com/AssetMantle/modules/x/orders/transactions/cancel"
	"github.com/AssetMantle/modules/x/orders/transactions/define"
	"github.com/AssetMantle/modules/x/orders/transactions/deputize"
	"github.com/AssetMantle/modules/x/orders/transactions/get"
	"github.com/AssetMantle/modules/x/orders/transactions/immediate"
	"github.com/AssetMantle/modules/x/orders/transactions/make"
	"github.com/AssetMantle/modules/x/orders/transactions/modify"
	"github.com/AssetMantle/modules/x/orders/transactions/put"
	"github.com/AssetMantle/modules/x/orders/transactions/revoke"
	"github.com/AssetMantle/modules/x/orders/transactions/take"
)

func (simulator) WeightedOperations(simulationState module.SimulationState, module helpers.Module) simulation.WeightedOperations {
	var weightMsg int

	simulationState.AppParams.GetOrGenerate(OpWeightMsg, &weightMsg, nil,
		func(_ *rand.Rand) {
			weightMsg = DefaultWeightMsg
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateDefineMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateMakeMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateCancelMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateTakeMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateDeputizeAndRevokeMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateGetMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateImmediateMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateModifyMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulatePutMsg(module)),
		),
	}
}

func simulateDefineMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result

		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "define", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)
		message := GenerateDefineMessage(from.Address, fromID.(ids.IdentityID), rand)

		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data)), nil, nil
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
			return simulationTypes.NewOperationMsg(&make.Message{}, false, "error in make message"), nil, nil
		}

		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data)), nil, nil
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
			return simulationTypes.NewOperationMsg(&make.Message{}, false, "error in make message"), nil, nil
		}

		result, err = simulationModules.ExecuteMessage(context, module, makeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(makeMessage, false, err.Error()), nil, nil
		}

		orderID := baseIDs.NewOrderID(makeMessage.(*make.Message).ClassificationID, baseQualified.NewImmutables(makeMessage.(*make.Message).ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(makeMessage.(*make.Message).ImmutableProperties.Get()...)...)))

		cancelMessage := cancel.NewMessage(from.Address, makeMessage.(*make.Message).FromID, orderID)
		result, err = simulationModules.ExecuteMessage(context, module, cancelMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(cancelMessage, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(cancelMessage, true, string(result.Data)), nil, nil
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
			return simulationTypes.NewOperationMsg(&make.Message{}, false, "error in make message"), nil, nil
		}

		result, err = simulationModules.ExecuteMessage(context, module, makeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(makeMessage, false, err.Error()), nil, nil
		}

		orderID := baseIDs.NewOrderID(makeMessage.(*make.Message).ClassificationID, baseQualified.NewImmutables(makeMessage.(*make.Message).ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(makeMessage.(*make.Message).ImmutableProperties.Get()...)...)))

		takeMessage := take.NewMessage(to.Address, makeMessage.(*make.Message).TakerID, math.NewInt(1), orderID)
		result, err = simulationModules.ExecuteMessage(context, module, takeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(takeMessage, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(takeMessage, true, string(result.Data)), nil, nil
	}
}

func GetMakeMessage(from, to simulationTypes.Account, rand *rand.Rand) sdkTypes.Msg {
	identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
	if err != nil {
		return nil
	}
	fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

	toIdentityIDString, err := simulationModules.LookupIdentityID(to.Address.String())
	if err != nil {
		return nil
	}
	toID, _ := baseIDs.PrototypeIdentityID().FromString(toIdentityIDString)

	classificationIDString, err := simulationModules.LookupOrderClassificationID(from.Address.String())
	if err != nil {
		return nil
	}
	classificationID, _ := baseIDs.PrototypeClassificationID().FromString(classificationIDString)

	assetIDString, err := simulationModules.LookupAssetID(from.Address.String())
	if err != nil {
		return nil
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
	// Regenerate immutable property values with same keys but new random data to
	// produce unique order IDs while conforming to the classification's structure.
	for _, i := range mappable.GetOrder().Get().GetImmutables().GetImmutablePropertyList().Get() {
		if i.IsMeta() {
			immutableMetaProperties = immutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
		} else {
			immutableProperties = immutableProperties.Add(baseProperties.NewMesaProperty(i.Get().GetKey(), baseTypes.GenerateRandomData(rand, rand.Intn(8)))).(*baseLists.PropertyList)
		}
	}
	for _, i := range mappable.GetOrder().Get().GetMutables().GetMutablePropertyList().Get() {
		if i.IsMeta() {
			mutableMetaProperties = mutableMetaProperties.Add(i).(*baseLists.PropertyList)
		} else {
			mutableProperties = mutableProperties.Add(i).(*baseLists.PropertyList)
		}
	}

	return make.NewMessage(from.Address, fromID.(ids.IdentityID), classificationID.(ids.ClassificationID), toID.(ids.IdentityID), assetID.(ids.AssetID), baseDocuments.NewCoinAsset("stake").GetCoinAssetID(), baseTypesGo.NewHeight(int64(rand.Intn(100)+10)), math.NewInt(1), math.NewInt(1), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
}
func simulateDeputizeAndRevokeMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "deputize", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		toIdentityIDString, err := simulationModules.LookupIdentityID(to.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "deputize", "no identity data"), nil, nil
		}
		toID, _ := baseIDs.PrototypeIdentityID().FromString(toIdentityIDString)

		classificationIDString, err := simulationModules.LookupOrderClassificationID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "deputize", "no order data"), nil, nil
		}
		classificationID, _ := baseIDs.PrototypeClassificationID().FromString(classificationIDString)

		orderMappable := &mappable.Mappable{}
		base.CodecPrototype().Unmarshal(orders.GetMappableBytes(classificationIDString), orderMappable)
		if orderMappable.Order == nil {
			return simulationTypes.NewOperationMsg(&deputize.Message{}, false, "nil order"), nil, nil
		}

		deputizeMessage := deputize.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), classificationID.(ids.ClassificationID), orderMappable.Order.Mutables.PropertyList, true, true, true, true, true)
		_, err = simulationModules.ExecuteMessage(context, module, deputizeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(deputizeMessage, false, err.Error()), nil, nil
		}

		revokeMessage := revoke.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), classificationID.(ids.ClassificationID))
		result, err := simulationModules.ExecuteMessage(context, module, revokeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(revokeMessage, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(revokeMessage, true, string(result.Data)), nil, nil
	}
}
func simulateGetMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		makeMessage := GetMakeMessage(from, to, rand)
		if makeMessage == nil {
			return simulationTypes.NewOperationMsg(&make.Message{}, false, "error in make message"), nil, nil
		}

		_, err := simulationModules.ExecuteMessage(context, module, makeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(makeMessage, false, err.Error()), nil, nil
		}

		orderID := baseIDs.NewOrderID(makeMessage.(*make.Message).ClassificationID, baseQualified.NewImmutables(makeMessage.(*make.Message).ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(makeMessage.(*make.Message).ImmutableProperties.Get()...)...)))

		// The taker identity (TakerID) belongs to the `to` account, so use
		// to.Address as the signer and TakerID as the fromID for authentication.
		getMessage := get.NewMessage(to.Address, makeMessage.(*make.Message).TakerID, orderID)
		result, err := simulationModules.ExecuteMessage(context, module, getMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(getMessage, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(getMessage, true, string(result.Data)), nil, nil
	}
}
func simulateImmediateMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		message := GetImmediateMessage(from, to, rand)
		if message == nil {
			return simulationTypes.NewOperationMsg(&immediate.Message{}, false, "error in immediate message"), nil, nil
		}

		result, err := simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data)), nil, nil
	}
}
func simulateModifyMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		makeMessage := GetMakeMessage(from, to, rand)
		if makeMessage == nil {
			return simulationTypes.NewOperationMsg(&make.Message{}, false, "error in make message"), nil, nil
		}

		_, err := simulationModules.ExecuteMessage(context, module, makeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(makeMessage, false, err.Error()), nil, nil
		}

		orderID := baseIDs.NewOrderID(makeMessage.(*make.Message).ClassificationID, baseQualified.NewImmutables(makeMessage.(*make.Message).ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(makeMessage.(*make.Message).ImmutableProperties.Get()...)...)))

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "modify", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		modifyMessage := modify.NewMessage(from.Address, fromID.(ids.IdentityID), orderID, math.NewInt(1), math.NewInt(1), baseTypesGo.NewHeight(int64(rand.Intn(100)+10)), makeMessage.(*make.Message).MutableMetaProperties, makeMessage.(*make.Message).MutableProperties)
		result, err := simulationModules.ExecuteMessage(context, module, modifyMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(modifyMessage, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(modifyMessage, true, string(result.Data)), nil, nil
	}
}
func simulatePutMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "put", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		assetIDString, err := simulationModules.LookupAssetID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "put", "no asset data"), nil, nil
		}
		assetID, _ := baseIDs.PrototypeAssetID().FromString(assetIDString)

		expiryHeight := baseTypesGo.NewHeight(context.BlockHeight() + int64(rand.Intn(1000)+100))
		message := put.NewMessage(from.Address, fromID.(ids.IdentityID), assetID.(ids.AssetID), baseDocuments.NewCoinAsset("stake").GetCoinAssetID(), math.NewInt(1), math.NewInt(1), expiryHeight)
		result, err := simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data)), nil, nil
	}
}

func GetImmediateMessage(from, to simulationTypes.Account, rand *rand.Rand) sdkTypes.Msg {
	identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
	if err != nil {
		return nil
	}
	fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

	toIdentityIDString, err := simulationModules.LookupIdentityID(to.Address.String())
	if err != nil {
		return nil
	}
	toID, _ := baseIDs.PrototypeIdentityID().FromString(toIdentityIDString)

	classificationIDString, err := simulationModules.LookupOrderClassificationID(from.Address.String())
	if err != nil {
		return nil
	}
	classificationID, _ := baseIDs.PrototypeClassificationID().FromString(classificationIDString)

	assetIDString, err := simulationModules.LookupAssetID(from.Address.String())
	if err != nil {
		return nil
	}
	assetID, _ := baseIDs.PrototypeAssetID().FromString(assetIDString)

	orderMappable := &mappable.Mappable{}
	base.CodecPrototype().Unmarshal(orders.GetMappableBytes(classificationIDString), orderMappable)
	immutableMetaProperties := &baseLists.PropertyList{}
	immutableProperties := &baseLists.PropertyList{}
	mutableMetaProperties := &baseLists.PropertyList{}
	mutableProperties := &baseLists.PropertyList{}
	if orderMappable.Order == nil {
		return nil
	}
	for _, i := range orderMappable.GetOrder().Get().GetImmutables().GetImmutablePropertyList().Get() {
		if i.IsMeta() {
			immutableMetaProperties = immutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
		} else {
			immutableProperties = immutableProperties.Add(baseProperties.NewMesaProperty(i.Get().GetKey(), baseTypes.GenerateRandomData(rand, rand.Intn(8)))).(*baseLists.PropertyList)
		}
	}
	for _, i := range orderMappable.GetOrder().Get().GetMutables().GetMutablePropertyList().Get() {
		if i.IsMeta() {
			mutableMetaProperties = mutableMetaProperties.Add(i).(*baseLists.PropertyList)
		} else {
			mutableProperties = mutableProperties.Add(i).(*baseLists.PropertyList)
		}
	}

	return immediate.NewMessage(from.Address, fromID.(ids.IdentityID), classificationID.(ids.ClassificationID), toID.(ids.IdentityID), assetID.(ids.AssetID), baseDocuments.NewCoinAsset("stake").GetCoinAssetID(), baseTypesGo.NewHeight(int64(rand.Intn(100)+10)), math.NewInt(1), math.NewInt(1), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
}

func GenerateDefineMessage(from sdkTypes.AccAddress, identityID ids.IdentityID, r *rand.Rand) helpers.Message {
	return define.NewMessage(from, identityID, baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r), baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r)).(helpers.Message)
}
