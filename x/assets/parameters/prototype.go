// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/parameters/burn_enabled"
	"github.com/AssetMantle/modules/x/assets/parameters/mint_enabled"
	"github.com/AssetMantle/modules/x/assets/parameters/renumerate_enabled"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(constants.ModuleName, burn_enabled.ValidatableParameter, mint_enabled.ValidatableParameter, renumerate_enabled.ValidatableParameter)
}
