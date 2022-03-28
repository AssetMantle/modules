/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Module interface {
	sdkTypesModule.AppModuleBasic
	sdkTypesModule.AppModule
	sdkTypesModule.AppModuleSimulation

	GetAuxiliary(string) Auxiliary

	DecodeModuleTransactionRequest(string, json.RawMessage) (sdkTypes.Msg, error)

	Initialize(*sdkTypes.KVStoreKey, paramTypes.Subspace, ...interface{}) Module
}
