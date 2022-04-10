// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
)

type Module interface {
	sdkTypesModule.AppModuleBasic
	sdkTypesModule.AppModule
	sdkTypesModule.AppModuleSimulation

	GetAuxiliary(string) Auxiliary

	DecodeModuleTransactionRequest(string, json.RawMessage) (sdkTypes.Msg, error)

	Initialize(*sdkTypes.KVStoreKey, params.Subspace, ...interface{}) Module
}
