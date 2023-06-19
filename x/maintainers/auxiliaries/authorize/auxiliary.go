// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authorize

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/utilities/name"
)

type dummy struct{}

var Auxiliary = baseHelpers.NewAuxiliary(
	name.GetPackageName(dummy{}),

	keeperPrototype,
)
