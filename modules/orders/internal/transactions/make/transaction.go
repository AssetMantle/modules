// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var Transaction = baseHelpers.NewTransaction(
	"make",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,
	constants.ClassificationID,
	constants.FromID,
	constants.MakerOwnableSplit,
	constants.MakerOwnableID,
	constants.TakerOwnableSplit,
	constants.TakerOwnableID,
	constants.ExpiresIn,
	constants.ImmutableMetaProperties,
	constants.ImmutableProperties,
	constants.MutableMetaProperties,
	constants.MutableProperties,
)
