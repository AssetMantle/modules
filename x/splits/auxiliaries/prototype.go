// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/burn"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/purge"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

func Prototype() helpers.Auxiliaries {
	return baseHelpers.NewAuxiliaries(
		burn.Auxiliary,
		mint.Auxiliary,
		purge.Auxiliary,
		renumerate.Auxiliary,
		transfer.Auxiliary,
	)
}
