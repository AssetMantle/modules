// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/deputize"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/revoke"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/super"
)

func Prototype() helpers.Auxiliaries {
	return baseHelpers.NewAuxiliaries(
		deputize.Auxiliary,
		maintain.Auxiliary,
		revoke.Auxiliary,
		super.Auxiliary,
		authorize.Auxiliary,
	)
}
