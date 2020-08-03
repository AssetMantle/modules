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
	InitializeKeepers(...interface{})
}
