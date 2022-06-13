// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var Transaction = baseHelpers.NewTransaction(
	"burn",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	constants.FromID,
	constants.AssetID,
)
