// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/modules/identities/internal/module"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

var Query = baseHelpers.NewQuery(
	"identities",
	"",
	"",

	module.Name,

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	flags.IdentityID,
)
