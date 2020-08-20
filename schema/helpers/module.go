/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
)

type Module interface {
	sdkTypesModule.AppModuleBasic
	sdkTypesModule.AppModule

	GetKVStoreKey() *sdkTypes.KVStoreKey
	GetDefaultParamspace() string
	GetAuxiliary(string) Auxiliary
	Initialize(...interface{}) Module
}
