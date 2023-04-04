// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/x/classifications/internal/module"
	"github.com/AssetMantle/modules/x/classifications/internal/parameters/bondRate"
	"github.com/AssetMantle/modules/x/classifications/internal/parameters/maxPropertyCount"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(module.Name, bondRate.ValidatableParameter, maxPropertyCount.ValidatableParameter)
}
