// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var Transaction = baseHelpers.NewTransaction(
	"modify",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,
	constants.FromID,
	constants.OrderID,
	constants.MakerOwnableSplit,
	constants.TakerOwnableSplit,
	constants.ExpiresIn,
	constants.ImmutableMetaProperties,
	constants.ImmutableProperties,
	constants.MutableMetaProperties,
	constants.MutableProperties,
)
