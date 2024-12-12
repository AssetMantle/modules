// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/utilities/name"
)

type dummy struct{}

var Auxiliary = baseHelpers.NewAuxiliary(
	name.GetPackageName(dummy{}),

	keeperPrototype,
)
