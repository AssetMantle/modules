// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"cosmossdk.io/math"
	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
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
	constantProperties "github.com/AssetMantle/schema/properties/constants"
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
			simulationModules.SafeOperation("orders/define", simulateDefineMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation("orders/make", simulateMakeMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation("orders/cancel", simulateCancelMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation("orders/take", simulateTakeMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation("orders/deputizeAndRevoke", simulateDeputizeAndRevokeMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation("orders/get", simulateGetMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation("orders/immediate", simulateImmediateMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation("orders/modify", simulateModifyMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation("orders/put", simulatePutMsg(module)),
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
			return simulationModules.RejectionOrError("orders", "define", err)
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data)), nil, nil
	}
}
func simulateMakeMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		makeMsg, _, err := DefineAndMake(context, module, from, to, rand)
		if err != nil {
			return simulationModules.RejectionOrError("orders", "make", err)
		}
		return simulationTypes.NewOperationMsg(makeMsg, true, ""), nil, nil
	}
}
func simulateCancelMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		makeMsg, orderID, err := DefineAndMake(context, module, from, to, rand)
		if err != nil {
			return simulationModules.RejectionOrError("orders", "cancel", err)
		}

		cancelMessage := cancel.NewMessage(from.Address, makeMsg.(*make.Message).FromID, orderID)
		result, err := simulationModules.ExecuteMessage(context, module, cancelMessage.(helpers.Message))
		if err != nil {
			return simulationModules.RejectionOrError("orders", "cancel", err)
		}
		return simulationTypes.NewOperationMsg(cancelMessage, true, string(result.Data)), nil, nil
	}
}
func simulateTakeMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		makeMsg, orderID, err := DefineAndMake(context, module, from, to, rand)
		if err != nil {
			return simulationModules.RejectionOrError("orders", "take", err)
		}

		takeMessage := take.NewMessage(to.Address, makeMsg.(*make.Message).TakerID, math.NewInt(1), orderID)
		result, err := simulationModules.ExecuteMessage(context, module, takeMessage.(helpers.Message))
		if err != nil {
			return simulationModules.RejectionOrError("orders", "take", err)
		}
		return simulationTypes.NewOperationMsg(takeMessage, true, string(result.Data)), nil, nil
	}
}

// DefineAndMake defines a new order classification and creates an order under it.
// Returns the make message and the order ID, or the failure's error.
func DefineAndMake(context sdkTypes.Context, module helpers.Module, from, to simulationTypes.Account, rand *rand.Rand) (sdkTypes.Msg, ids.OrderID, error) {
	identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
	if err != nil {
		return nil, nil, err
	}
	fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

	toIdentityIDString, err := simulationModules.LookupIdentityID(to.Address.String())
	if err != nil {
		return nil, nil, err
	}
	toID, _ := baseIDs.PrototypeIdentityID().FromString(toIdentityIDString)

	assetIDString, err := simulationModules.LookupAssetID(from.Address.String())
	if err != nil {
		return nil, nil, err
	}
	assetID, _ := baseIDs.PrototypeAssetID().FromString(assetIDString)

	// IMPORTANT: PropertyList.Add mutates in place. Snapshot all random
	// properties to prevent shared-state corruption between steps.
	immMetaSnap := baseTypes.GenerateRandomMetaPropertyList(rand).Get()
	immSnap := baseTypes.GenerateRandomPropertyList(rand).Get()
	mutMetaSnap := baseTypes.GenerateRandomMetaPropertyList(rand).Get()
	mutSnap := baseTypes.GenerateRandomPropertyList(rand).Get()

	// TakerSplitProperty must reveal a non-zero value: order keeper's get path
	// transfers `order.GetTakerSplit()` of takerAssetID from taker to maker.
	// Prototype's zero data causes the transfer aux to error with
	// "value must be greater than zero".
	// BondAmountProperty: the make keeper requires it revealed in mutables, and
	// its value must cover rate*totalWeight for any genesis bond rate; the
	// orders define keeper adds the order constants below before weighing.
	revealedTakerSplit := baseProperties.NewMetaProperty(constantProperties.TakerSplitProperty.GetKey(), baseData.NewNumberData(math.NewInt(1)))
	bondProperty := baseProperties.NewMetaProperty(
		constantProperties.BondAmountProperty.GetKey(),
		baseData.NewNumberData(simulationModules.SumBondWeights(
			[][]properties.AnyProperty{immMetaSnap, immSnap, mutMetaSnap, mutSnap},
			constantProperties.TakerSplitProperty, constantProperties.BondAmountProperty,
			constantProperties.ExchangeRateProperty, constantProperties.CreationHeightProperty,
			constantProperties.MakerAssetIDProperty, constantProperties.TakerAssetIDProperty,
			constantProperties.MakerIDProperty, constantProperties.TakerIDProperty,
			constantProperties.ExpiryHeightProperty, constantProperties.MakerSplitProperty,
		).MulRaw(simulationModules.MaxBondRate)))
	defineMessage := define.NewMessage(from.Address, fromID.(ids.IdentityID),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immSnap...)...),
		baseLists.NewPropertyList(revealedTakerSplit, bondProperty).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(mutSnap...)...))
	_, err = simulationModules.ExecuteMessage(context, module, defineMessage.(helpers.Message))
	if err != nil {
		return nil, nil, err
	}

	// Compute classificationID with keeper-added properties (fresh lists).
	immutables := baseQualified.NewImmutables(
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...).Add(baseLists.AnyPropertiesToProperties(immSnap...)...).Add(
			constantProperties.ExchangeRateProperty, constantProperties.CreationHeightProperty,
			constantProperties.MakerAssetIDProperty, constantProperties.TakerAssetIDProperty,
			constantProperties.MakerIDProperty, constantProperties.TakerIDProperty))
	mutables := baseQualified.NewMutables(
		baseLists.NewPropertyList(revealedTakerSplit, bondProperty).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...).Add(baseLists.AnyPropertiesToProperties(mutSnap...)...).Add(
			constantProperties.ExpiryHeightProperty, constantProperties.MakerSplitProperty))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)

	// Make message with fresh lists, TakerSplit + BondAmount in mutable meta
	makeMsg := make.NewMessage(from.Address, fromID.(ids.IdentityID), classificationID, toID.(ids.IdentityID), assetID.(ids.AssetID),
		baseDocuments.NewCoinAsset("stake").GetCoinAssetID(), baseTypesGo.NewHeight(int64(rand.Intn(100)+10)), math.NewInt(1), math.NewInt(1),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immSnap...)...),
		baseLists.NewPropertyList(revealedTakerSplit, bondProperty).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(mutSnap...)...))
	result, err := simulationModules.ExecuteMessage(context, module, makeMsg.(helpers.Message))
	if err != nil {
		return nil, nil, err
	}

	// Parse orderID from response (keeper computes it from runtime data values
	// like ExchangeRate/CreationHeight; local immutables don't have those).
	orderIDProto, _ := baseIDs.PrototypeOrderID().FromString(string(result.Data))
	orderID, _ := orderIDProto.(ids.OrderID)
	return makeMsg, orderID, nil
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
			mutableMetaProperties = mutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
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

		// Define a new classification (creates maintainer with super permissions),
		// then deputize the `to` identity under it.
		makeMsg, _, err := DefineAndMake(context, module, from, to, rand)
		if err != nil {
			return simulationModules.RejectionOrError("orders", "deputize", err)
		}

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

		classificationID := makeMsg.(*make.Message).ClassificationID

		// Pass empty maintained properties — super maintainer's mutable list
		// excludes BondAmount/Authentication (super aux strips them), so
		// passing message.MutableMetaProperties (which includes BondAmount)
		// would fail the MaintainsProperty check in maintainer/deputize aux.
		deputizeMessage := deputize.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), classificationID, baseLists.NewPropertyList(), true, true, true, true, true)
		_, err = simulationModules.ExecuteMessage(context, module, deputizeMessage.(helpers.Message))
		if err != nil {
			return simulationModules.RejectionOrError("orders", "deputize", err)
		}

		revokeMessage := revoke.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), classificationID)
		result, err := simulationModules.ExecuteMessage(context, module, revokeMessage.(helpers.Message))
		if err != nil {
			return simulationModules.RejectionOrError("orders", "revoke", err)
		}
		return simulationTypes.NewOperationMsg(revokeMessage, true, string(result.Data)), nil, nil
	}
}
func simulateGetMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		makeMsg, orderID, err := DefineAndMake(context, module, from, to, rand)
		if err != nil {
			return simulationModules.RejectionOrError("orders", "get", err)
		}

		getMessage := get.NewMessage(to.Address, makeMsg.(*make.Message).TakerID, orderID)
		result, err := simulationModules.ExecuteMessage(context, module, getMessage.(helpers.Message))
		if err != nil {
			return simulationModules.RejectionOrError("orders", "get", err)
		}
		return simulationTypes.NewOperationMsg(getMessage, true, string(result.Data)), nil, nil
	}
}
func simulateImmediateMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		// Immediate uses the same define+make pattern but with immediate.NewMessage
		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "immediate", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		toIdentityIDString, err := simulationModules.LookupIdentityID(to.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "immediate", "no identity data"), nil, nil
		}
		toID, _ := baseIDs.PrototypeIdentityID().FromString(toIdentityIDString)

		assetIDString, err := simulationModules.LookupAssetID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "immediate", "no asset data"), nil, nil
		}
		assetID, _ := baseIDs.PrototypeAssetID().FromString(assetIDString)

		// Snapshot to prevent PropertyList.Add mutation across steps.
		immMetaSnap := baseTypes.GenerateRandomMetaPropertyList(rand).Get()
		immSnap := baseTypes.GenerateRandomPropertyList(rand).Get()
		mutMetaSnap := baseTypes.GenerateRandomMetaPropertyList(rand).Get()
		mutSnap := baseTypes.GenerateRandomPropertyList(rand).Get()

		// TakerSplit revealed with non-zero data — required for downstream
		// transfers; bond revealed covering rate*totalWeight for any genesis
		// bond rate; see DefineAndMake for full explanation.
		revealedTakerSplit := baseProperties.NewMetaProperty(constantProperties.TakerSplitProperty.GetKey(), baseData.NewNumberData(math.NewInt(1)))
		bondProperty := baseProperties.NewMetaProperty(
			constantProperties.BondAmountProperty.GetKey(),
			baseData.NewNumberData(simulationModules.SumBondWeights(
				[][]properties.AnyProperty{immMetaSnap, immSnap, mutMetaSnap, mutSnap},
				constantProperties.TakerSplitProperty, constantProperties.BondAmountProperty,
				constantProperties.ExchangeRateProperty, constantProperties.CreationHeightProperty,
				constantProperties.MakerAssetIDProperty, constantProperties.TakerAssetIDProperty,
				constantProperties.MakerIDProperty, constantProperties.TakerIDProperty,
				constantProperties.ExpiryHeightProperty, constantProperties.MakerSplitProperty,
			).MulRaw(simulationModules.MaxBondRate)))
		defineMessage := define.NewMessage(from.Address, fromID.(ids.IdentityID),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immSnap...)...),
			baseLists.NewPropertyList(revealedTakerSplit, bondProperty).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(mutSnap...)...))
		_, err = simulationModules.ExecuteMessage(context, module, defineMessage.(helpers.Message))
		if err != nil {
			return simulationModules.RejectionOrError("orders", "immediate", err)
		}

		immutables := baseQualified.NewImmutables(
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...).Add(baseLists.AnyPropertiesToProperties(immSnap...)...).Add(
				constantProperties.ExchangeRateProperty, constantProperties.CreationHeightProperty,
				constantProperties.MakerAssetIDProperty, constantProperties.TakerAssetIDProperty,
				constantProperties.MakerIDProperty, constantProperties.TakerIDProperty))
		mutables := baseQualified.NewMutables(
			baseLists.NewPropertyList(revealedTakerSplit, bondProperty).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...).Add(baseLists.AnyPropertiesToProperties(mutSnap...)...).Add(
				constantProperties.ExpiryHeightProperty, constantProperties.MakerSplitProperty))
		classificationID := baseIDs.NewClassificationID(immutables, mutables)

		immediateMsg := immediate.NewMessage(from.Address, fromID.(ids.IdentityID), classificationID, toID.(ids.IdentityID), assetID.(ids.AssetID),
			baseDocuments.NewCoinAsset("stake").GetCoinAssetID(), baseTypesGo.NewHeight(int64(rand.Intn(100)+10)), math.NewInt(1), math.NewInt(1),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immSnap...)...),
			baseLists.NewPropertyList(revealedTakerSplit, bondProperty).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(mutSnap...)...))
		result, err := simulationModules.ExecuteMessage(context, module, immediateMsg.(helpers.Message))
		if err != nil {
			return simulationModules.RejectionOrError("orders", "immediate", err)
		}
		return simulationTypes.NewOperationMsg(immediateMsg, true, string(result.Data)), nil, nil
	}
}
func simulateModifyMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		makeMsg, orderID, err := DefineAndMake(context, module, from, to, rand)
		if err != nil {
			return simulationModules.RejectionOrError("orders", "modify", err)
		}

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("orders", "modify", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		modifyMessage := modify.NewMessage(from.Address, fromID.(ids.IdentityID), orderID, math.NewInt(1), math.NewInt(1), baseTypesGo.NewHeight(int64(rand.Intn(100)+10)), makeMsg.(*make.Message).MutableMetaProperties, makeMsg.(*make.Message).MutableProperties)
		result, err := simulationModules.ExecuteMessage(context, module, modifyMessage.(helpers.Message))
		if err != nil {
			return simulationModules.RejectionOrError("orders", "modify", err)
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

		expiryHeight := baseTypesGo.NewHeight(context.BlockHeight() + int64(rand.Intn(50)+10))
		message := put.NewMessage(from.Address, fromID.(ids.IdentityID), assetID.(ids.AssetID), baseDocuments.NewCoinAsset("stake").GetCoinAssetID(), math.NewInt(1), math.NewInt(1), expiryHeight)
		result, err := simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationModules.RejectionOrError("orders", "put", err)
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
			mutableMetaProperties = mutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
		} else {
			mutableProperties = mutableProperties.Add(i).(*baseLists.PropertyList)
		}
	}

	return immediate.NewMessage(from.Address, fromID.(ids.IdentityID), classificationID.(ids.ClassificationID), toID.(ids.IdentityID), assetID.(ids.AssetID), baseDocuments.NewCoinAsset("stake").GetCoinAssetID(), baseTypesGo.NewHeight(int64(rand.Intn(100)+10)), math.NewInt(1), math.NewInt(1), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
}

func GenerateDefineMessage(from sdkTypes.AccAddress, identityID ids.IdentityID, r *rand.Rand) helpers.Message {
	return define.NewMessage(from, identityID, baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r), baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r)).(helpers.Message)
}
