// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/x/assets/internal/module"
	"github.com/AssetMantle/modules/x/assets/internal/parameters/burnEnabled"
	"github.com/AssetMantle/modules/x/assets/internal/parameters/mintEnabled"
	"github.com/AssetMantle/modules/x/assets/internal/parameters/renumerateEnabled"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(module.Name, burnEnabled.ValidatableParameter, mintEnabled.ValidatableParameter, renumerateEnabled.ValidatableParameter)
}
