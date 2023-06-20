// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/constants"
	"github.com/AssetMantle/modules/x/splits/parameters/wrap_allowed_coins"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(constants.ModuleName, wrap_allowed_coins.ValidatableParameter)
}
