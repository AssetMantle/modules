// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"make",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,
	flags.ClassificationID,
	flags.FromID,
	flags.MakerOwnableSplit,
	flags.MakerOwnableID,
	flags.TakerOwnableSplit,
	flags.TakerOwnableID,
	flags.ExpiresIn,
	flags.ImmutableMetaProperties,
	flags.ImmutableProperties,
	flags.MutableMetaProperties,
	flags.MutableProperties,
)
