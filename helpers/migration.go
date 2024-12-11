package helpers

import (
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Migration interface {
	GetHandler() sdkModuleTypes.MigrationHandler
	Initialize(Mapper, ParameterManager, paramsTypes.Subspace) Migration
}
