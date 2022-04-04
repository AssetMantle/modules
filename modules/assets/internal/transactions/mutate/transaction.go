// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"mutate",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	flags.AssetID,
	flags.FromID,
	flags.MutableMetaProperties,
	flags.MutableProperties,
)
