// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
)

func Prototype() helpers.Auxiliaries {
	return baseHelpers.NewAuxiliaries(
		authenticate.Auxiliary,
	)
}
