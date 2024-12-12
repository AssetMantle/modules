// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/parameters/burn_enabled"
	"github.com/AssetMantle/modules/x/assets/parameters/mint_enabled"
	"github.com/AssetMantle/modules/x/assets/parameters/renumerate_enabled"
	"github.com/AssetMantle/modules/x/assets/parameters/unwrap_allowed_coins"
	"github.com/AssetMantle/modules/x/assets/parameters/wrap_allowed_coins"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(burn_enabled.ValidatableParameter, mint_enabled.ValidatableParameter, renumerate_enabled.ValidatableParameter, unwrap_allowed_coins.ValidatableParameter, wrap_allowed_coins.ValidatableParameter)
}
