// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"encoding/json"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Module interface {
	BasicModule
	sdkModuleTypes.AppModuleGenesis
	sdkModuleTypes.EndBlockAppModule
	sdkModuleTypes.BeginBlockAppModule

	sdkModuleTypes.HasConsensusVersion
	sdkModuleTypes.HasInvariants
	sdkModuleTypes.HasServices

	sdkModuleTypes.AppModuleSimulation
	sdkModuleTypes.HasProposalContents
	sdkModuleTypes.HasProposalMsgs

	GetAuxiliary(string) Auxiliary

	DecodeModuleTransactionRequest(string, json.RawMessage) (sdkTypes.Msg, error)

	Initialize(*storeTypes.KVStoreKey, paramsTypes.Subspace, ...interface{}) Module

	GetTransactions() Transactions
}
