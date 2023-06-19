// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/utilities/name"
)

type dummy struct{}

var ModuleName = name.GetSuperPackageName(dummy{})

const ModuleConsensusVersion = 1

var ModuleStoreKeyPrefix = constants.MetasStoreKeyPrefix
