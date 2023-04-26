// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/module"
	"github.com/AssetMantle/modules/x/assets/parameters/burnEnabled"
	"github.com/AssetMantle/modules/x/assets/parameters/mintEnabled"
	"github.com/AssetMantle/modules/x/assets/parameters/renumerateEnabled"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(module.Name, burnEnabled.ValidatableParameter, mintEnabled.ValidatableParameter, renumerateEnabled.ValidatableParameter)
}
