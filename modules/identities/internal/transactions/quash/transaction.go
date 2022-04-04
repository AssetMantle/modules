// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package quash

import (
	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"quash",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	flags.FromID,
	flags.IdentityID,
)
