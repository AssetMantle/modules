// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
)

func Prototype() helpers.Auxiliaries {
	return baseHelpers.NewAuxiliaries(
		scrub.Auxiliary,
		supplement.Auxiliary,
	)
}
