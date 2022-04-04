// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"burn",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	flags.FromID,
	flags.AssetID,
)
