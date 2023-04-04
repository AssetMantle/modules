// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/x/splits/auxiliaries/burn"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

func Prototype() helpers.Auxiliaries {
	return baseHelpers.NewAuxiliaries(
		burn.Auxiliary,
		mint.Auxiliary,
		renumerate.Auxiliary,
		transfer.Auxiliary,
	)
}
