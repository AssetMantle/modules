// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	"github.com/AssetMantle/modules/utilities/name"
	"github.com/AssetMantle/schema/documents/base"
)

type dummy struct{}

var ModuleName = name.GetSuperPackageName(dummy{})

const ModuleConsensusVersion = 2

var ModuleIdentity = base.NewModuleIdentity(ModuleName)
