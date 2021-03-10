/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"encoding/json"

	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
)

type Module interface {
	sdkTypesModule.AppModuleBasic
	sdkTypesModule.AppModule
	sdkTypesModule.AppModuleSimulation

	GetAuxiliary(string) Auxiliary

	DecodeModuleTransactionRequest(string, json.RawMessage) (sdkTypes.Msg, error)

	Initialize(*sdkTypes.KVStoreKey, paramsTypes.Subspace, ...interface{}) Module
}
