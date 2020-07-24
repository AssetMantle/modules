package helpers

import (
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
)

type Module interface {
	sdkTypesModule.AppModuleBasic
	sdkTypesModule.AppModule

	GetStoreKey() string
	GetDefaultParamspace() string
	GetAuxiliaryKeepers(...string) []AuxiliaryKeeper
	InitializeKeepers(...interface{})
}
