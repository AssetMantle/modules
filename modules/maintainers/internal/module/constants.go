// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids/base"
)

const Name = "maintainers"

var StoreKeyPrefix = constants.MaintainersStoreKeyPrefix

// TODO impl through enums
var (
	Mint       = base.NewStringID("mint")
	Burn       = base.NewStringID("burn")
	Renumerate = base.NewStringID("renumerate")
	Add        = base.NewStringID("add")
	Remove     = base.NewStringID("remove")
	Mutate     = base.NewStringID("mutate")
)
