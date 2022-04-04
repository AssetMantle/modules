// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package provision

import (
	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"provision",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	flags.To,
	flags.IdentityID,
)
