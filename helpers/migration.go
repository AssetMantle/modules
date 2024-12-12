// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Migration interface {
	GetHandler() sdkModuleTypes.MigrationHandler
	Initialize(Mapper, ParameterManager, paramsTypes.Subspace) Migration
}
