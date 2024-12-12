// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/helpers"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type migration struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
	paramsSubspace   paramsTypes.Subspace
	migrate          func(sdkTypes.Context, helpers.Mapper, helpers.ParameterManager, paramsTypes.Subspace) error
}

var _ helpers.Migration = (*migration)(nil)

func (migration migration) GetHandler() sdkModuleTypes.MigrationHandler {
	return func(context sdkTypes.Context) error {
		return migration.migrate(context, migration.mapper, migration.parameterManager, migration.paramsSubspace)
	}
}
func (migration migration) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, paramsSubspace paramsTypes.Subspace) helpers.Migration {
	migration.mapper = mapper
	migration.parameterManager = parameterManager
	migration.paramsSubspace = paramsSubspace
	return migration
}

func NewMigration(migrate func(sdkTypes.Context, helpers.Mapper, helpers.ParameterManager, paramsTypes.Subspace) error) helpers.Migration {
	return migration{
		migrate: migrate,
	}
}
