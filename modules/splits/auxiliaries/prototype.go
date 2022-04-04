// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/burn"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/mint"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Auxiliaries {
	return base.NewAuxiliaries(
		burn.Auxiliary,
		mint.Auxiliary,
		renumerate.Auxiliary,
		transfer.Auxiliary,
	)
}
