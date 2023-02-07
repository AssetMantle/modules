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
	constants.ExpiresIn,
	constants.FromID,
	constants.MakerOwnableID,
	constants.MakerOwnableSplit,
	constants.MutableMetaProperties,
	constants.MutableProperties,
	constants.TakerID,
	constants.TakerOwnableSplit,
	constants.TakerOwnableID,
	constants.ImmutableMetaProperties,
	constants.ImmutableProperties,
)
