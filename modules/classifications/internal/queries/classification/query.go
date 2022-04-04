// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classification

import (
	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

var Query = base.NewQuery(
	"classifications",
	"",
	"",

	module.Name,

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	flags.ClassificationID,
)
