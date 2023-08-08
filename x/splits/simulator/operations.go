// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	simulationModules "github.com/AssetMantle/modules/simulation"
	"github.com/AssetMantle/modules/simulation/simulated_database/assets"
	"github.com/AssetMantle/modules/simulation/simulated_database/identities"
	"github.com/AssetMantle/modules/x/splits/transactions/send"
	"github.com/AssetMantle/modules/x/splits/transactions/unwrap"
	"github.com/AssetMantle/modules/x/splits/transactions/wrap"
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
			simulateSendMsg(module),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulateWrapAndUnwrapMsg(module),
		),
	}
}

func simulateSendMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		var assetIDString, identityIDString string

		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
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

		assetMap := assets.GetAssetData(from.Address.String())
		for _, id := range assetMap {
			assetIDString = id
			break
		}

		assetID, _ := baseIDs.PrototypeAssetID().FromString(assetIDString)
		message := send.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), assetID.(ids.AssetID), sdkTypes.NewInt(1))

		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
func simulateWrapAndUnwrapMsg(module helpers.Module) simulationTypes.Operation {
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

		wrapMessage := wrap.NewMessage(from.Address, fromID.(ids.IdentityID), sdkTypes.NewCoins(sdkTypes.NewCoin("stake", sdkTypes.NewInt(1))))

		result, err = simulationModules.ExecuteMessage(context, module, wrapMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(wrapMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}

		unwrapMessage := unwrap.NewMessage(from.Address, fromID.(ids.IdentityID), baseIDs.GenerateCoinAssetID("stake"), sdkTypes.NewInt(1))

		result, err = simulationModules.ExecuteMessage(context, module, unwrapMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(unwrapMessage, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(unwrapMessage, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
