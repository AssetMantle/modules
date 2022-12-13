// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var Transaction = baseHelpers.NewTransaction(
	"mutate",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	constants.IdentityID,
	constants.FromID,
	constants.MutableMetaProperties,
	constants.MutableProperties,
)
