// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/modules/splits/internal/parameters/dummy"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Parameters {
	return base.NewParameters(dummy.Parameter)
}
