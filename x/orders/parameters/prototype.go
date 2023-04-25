// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/orders/module"
	"github.com/AssetMantle/modules/x/orders/parameters/maxOrderLife"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(module.Name, maxOrderLife.ValidatableParameter)
}
