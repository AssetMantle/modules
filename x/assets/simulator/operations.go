// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"

	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	simulationModules "github.com/AssetMantle/modules/simulation"
	baseTypes "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/assets/transactions/burn"
	"github.com/AssetMantle/modules/x/assets/transactions/define"
	"github.com/AssetMantle/modules/x/assets/transactions/deputize"
	"github.com/AssetMantle/modules/x/assets/transactions/mint"
	"github.com/AssetMantle/modules/x/assets/transactions/mutate"
	"github.com/AssetMantle/modules/x/assets/transactions/renumerate"
	"github.com/AssetMantle/modules/x/assets/transactions/revoke"
	"github.com/AssetMantle/modules/x/assets/transactions/send"
	"github.com/AssetMantle/modules/x/assets/transactions/unwrap"
	"github.com/AssetMantle/modules/x/assets/transactions/wrap"

)

func (simulator) WeightedOperations(simulationState module.SimulationState, module helpers.Module) simulation.WeightedOperations {
	var weightMsg int

	simulationModules.SimTxConfig = simulationState.TxConfig

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
			simulationModules.SafeOperation(simulateMintMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateBurnMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateRenumerateMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateDeputizeAndRevokeMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateMutateMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateSendMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateWrapAndUnwrapMsg(module)),
		),
	}
}

func simulateDefineMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		account, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		identityIDString, err := simulationModules.LookupIdentityID(account.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "define", "no identity data"), nil, nil
		}
		identityID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)
		message := GenerateDefineMessage(account.Address, identityID.(ids.IdentityID), rand).(*define.Message)
		result, err := simulationModules.ExecuteMessage(context, module, message)
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data)), nil, nil
	}
}
func simulateMintMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		mintMsg, _ := DefineAndMint(context, module, from, to, rand)
		if mintMsg == nil {
			return simulationTypes.NoOpMsg("assets", "mint", "define+mint failed"), nil, nil
		}
		return simulationTypes.NewOperationMsg(mintMsg, true, ""), nil, nil
	}
}
func simulateRenumerateMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "renumerate", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		_, assetID := DefineAndMint(context, module, from, to, rand)
		if assetID == nil {
			return simulationTypes.NoOpMsg("assets", "renumerate", "define+mint failed"), nil, nil
		}

		renumerateMessage := renumerate.NewMessage(from.Address, fromID.(ids.IdentityID), assetID)
		result, err := simulationModules.ExecuteMessage(context, module, renumerateMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "renumerate", err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(renumerateMessage, true, string(result.Data)), nil, nil
	}
}
func simulateBurnMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "burn", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		_, assetID := DefineAndMint(context, module, from, from, rand)
		if assetID == nil {
			return simulationTypes.NoOpMsg("assets", "burn", "define+mint failed"), nil, nil
		}

		burnMessage := burn.NewMessage(from.Address, fromID.(ids.IdentityID), assetID)
		result, err := simulationModules.ExecuteMessage(context, module, burnMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "burn", err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(burnMessage, true, string(result.Data)), nil, nil
	}
}
func simulateDeputizeAndRevokeMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "deputize", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		toIdentityIDString, err := simulationModules.LookupIdentityID(to.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "deputize", "no identity data"), nil, nil
		}
		toID, _ := baseIDs.PrototypeIdentityID().FromString(toIdentityIDString)

		mintMsg, _ := DefineAndMint(context, module, from, to, rand)
		if mintMsg == nil {
			return simulationTypes.NoOpMsg("assets", "deputize", "define+mint failed"), nil, nil
		}

		mintMessage := mintMsg.(*mint.Message)
		classificationID := mintMessage.ClassificationID

		// Pass empty maintained properties — the deputize grants permissions
		// (canMint, canBurn, etc.) regardless of specific property access.
		deputizeMessage := deputize.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), classificationID, baseLists.NewPropertyList(), true, true, true, true, true, true)
		_, err = simulationModules.ExecuteMessage(context, module, deputizeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "deputize", err.Error()), nil, nil
		}

		revokeMessage := revoke.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), classificationID)
		result, err := simulationModules.ExecuteMessage(context, module, revokeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "revoke", err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(revokeMessage, true, string(result.Data)), nil, nil
	}
}
func simulateMutateMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "mutate", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		_, assetID := DefineAndMint(context, module, from, to, rand)
		if assetID == nil {
			return simulationTypes.NoOpMsg("assets", "mutate", "define+mint failed"), nil, nil
		}

		// Pass empty property lists — the mutate keeper validates authentication
		// and authorization without requiring specific property updates.
		mutateMessage := mutate.NewMessage(from.Address, fromID.(ids.IdentityID), assetID, baseLists.NewPropertyList(), baseLists.NewPropertyList())
		result, err := simulationModules.ExecuteMessage(context, module, mutateMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "mutate", err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(mutateMessage, true, string(result.Data)), nil, nil
	}
}

func GenerateDefineMessage(from sdkTypes.AccAddress, identityID ids.IdentityID, r *rand.Rand) helpers.Message {
	return define.NewMessage(from, identityID, baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r), baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r)).(helpers.Message)
}
// DefineAndMint defines a new classification and mints an asset under it.
// Returns the mint message and the minted asset ID, or nil on failure.
func DefineAndMint(context sdkTypes.Context, module helpers.Module, from, to simulationTypes.Account, rand *rand.Rand) (sdkTypes.Msg, ids.AssetID) {
	identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
	if err != nil {
		return nil, nil
	}
	fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

	toIdentityIDString, err := simulationModules.LookupIdentityID(to.Address.String())
	if err != nil {
		return nil, nil
	}
	toID, _ := baseIDs.PrototypeIdentityID().FromString(toIdentityIDString)

	immutableMetaProps := baseTypes.GenerateRandomMetaPropertyList(rand)
	immutableProps := baseTypes.GenerateRandomPropertyList(rand)
	mutableMetaProps := baseTypes.GenerateRandomMetaPropertyList(rand)
	mutableProps := baseTypes.GenerateRandomPropertyList(rand)

	defineMessage := define.NewMessage(from.Address, fromID.(ids.IdentityID), immutableMetaProps, immutableProps, mutableMetaProps, mutableProps)
	_, err = simulationModules.ExecuteMessage(context, module, defineMessage.(helpers.Message))
	if err != nil {
		return nil, nil
	}

	immutables := baseQualified.NewImmutables(immutableMetaProps.Add(baseLists.AnyPropertiesToProperties(immutableProps.Get()...)...))
	mutables := baseQualified.NewMutables(mutableMetaProps.Add(baseLists.AnyPropertiesToProperties(mutableProps.Get()...)...))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	assetID := baseIDs.NewAssetID(classificationID, immutables)

	mintMsg := mint.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), classificationID, immutableMetaProps, immutableProps, mutableMetaProps, mutableProps)
	_, err = simulationModules.ExecuteMessage(context, module, mintMsg.(helpers.Message))
	if err != nil {
		return nil, nil
	}

	return mintMsg, assetID
}

func simulateSendMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result

		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "send", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		toIdentityIDString, err := simulationModules.LookupIdentityID(to.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "send", "no identity data"), nil, nil
		}
		toID, _ := baseIDs.PrototypeIdentityID().FromString(toIdentityIDString)

		assetIDString, err := simulationModules.LookupAssetID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "send", "no asset data"), nil, nil
		}

		assetID, _ := baseIDs.PrototypeAssetID().FromString(assetIDString)
		message := send.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), assetID.(ids.AssetID), math.NewInt(1))

		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data)), nil, nil
	}
}
func simulateWrapAndUnwrapMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result

		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("assets", "wrap", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		wrapMessage := wrap.NewMessage(from.Address, fromID.(ids.IdentityID), sdkTypes.NewCoins(sdkTypes.NewCoin("stake", math.NewInt(1))))

		result, err = simulationModules.ExecuteMessage(context, module, wrapMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(wrapMessage, false, err.Error()), nil, nil
		}

		unwrapMessage := unwrap.NewMessage(from.Address, fromID.(ids.IdentityID), sdkTypes.NewCoins(sdkTypes.NewCoin("stake", math.NewInt(1))))

		result, err = simulationModules.ExecuteMessage(context, module, unwrapMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(unwrapMessage, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(unwrapMessage, true, string(result.Data)), nil, nil
	}
}
