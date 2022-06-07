// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var Transaction = baseHelpers.NewTransaction(
	"mint",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	constants.ToID,
	constants.FromID,
	constants.ClassificationID,
	constants.ImmutableMetaProperties,
	constants.ImmutableProperties,
	constants.MutableMetaProperties,
	constants.MutableProperties,
)
