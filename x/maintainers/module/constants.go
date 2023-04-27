// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	"github.com/AssetMantle/modules/helpers/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
)

const Name = "maintainers"
const ConsensusVersion = 1

var StoreKeyPrefix = constants.MaintainersStoreKeyPrefix

var (
	Add        = baseIDs.NewStringID("add")
	Remove     = baseIDs.NewStringID("remove")
	Mutate     = baseIDs.NewStringID("mutate")
	Mint       = baseIDs.NewStringID("mint")
	Burn       = baseIDs.NewStringID("burn")
	Renumerate = baseIDs.NewStringID("renumerate")
)
