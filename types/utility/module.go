package utility

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
)

type Module interface {
	sdkTypesModule.AppModuleBasic
	sdkTypesModule.AppModule

	GetStoreKey() string
	GetDefaultParamspace() string
	GetAuxiliaryKeepers(...string) []AuxiliaryKeeper
	InitializeKeepers(*codec.Codec, sdkTypes.StoreKey, params.Subspace, ...interface{})
}
