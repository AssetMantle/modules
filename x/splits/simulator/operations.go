// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	simulationModules "github.com/AssetMantle/modules/simulation"
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/assets"
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/identities"
	"github.com/AssetMantle/modules/x/splits/transactions/send"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
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
			simulateSendMsg(module),
		),
	}
	return nil
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
		fromID, _ := baseIDs.ReadIdentityID(identityIDString)

		toIDMap := identities.GetIDData(to.Address.String())
		for _, id := range toIDMap {
			identityIDString = id
			break
		}
		toID, _ := baseIDs.ReadIdentityID(identityIDString)

		assetMap := assets.GetAssetData(from.Address.String())
		for _, id := range assetMap {
			assetIDString = id
			break
		}

		assetID, _ := baseIDs.ReadAssetID(assetIDString)
		message := send.NewMessage(from.Address, fromID, toID, assetID, sdkTypes.NewInt(1))

		result, err = simulationModules.ExecuteMessage(context, module, message.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error(), base.CodecPrototype().GetProtoCodec()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data), base.CodecPrototype().GetProtoCodec()), nil, nil
	}
}
