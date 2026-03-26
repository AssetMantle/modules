// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	storeTypes "cosmossdk.io/store/types"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gorilla/mux"
)

type Module interface {
	BasicModule
	sdkModuleTypes.HasABCIGenesis

	sdkModuleTypes.HasConsensusVersion
	sdkModuleTypes.HasInvariants
	sdkModuleTypes.HasServices

	sdkModuleTypes.AppModuleSimulation

	GetAuxiliary(string) Auxiliary

	Initialize(*storeTypes.KVStoreKey, paramsTypes.Subspace, ...interface{}) Module

	GetTransactions() Transactions

	RegisterRESTRoutes(client.Context, *mux.Router)
}
