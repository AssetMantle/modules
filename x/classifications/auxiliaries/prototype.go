// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/bond"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/burn"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/unbond"
)

func Prototype() helpers.Auxiliaries {
	return baseHelpers.NewAuxiliaries(
		conform.Auxiliary,
		define.Auxiliary,
		member.Auxiliary,
		bond.Auxiliary,
		burn.Auxiliary,
		unbond.Auxiliary,
	)
}
