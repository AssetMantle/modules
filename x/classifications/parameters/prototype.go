// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/parameters/bond_rate"
	"github.com/AssetMantle/modules/x/classifications/parameters/define_enabled"
	"github.com/AssetMantle/modules/x/classifications/parameters/max_property_count"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(bond_rate.ValidatableParameter, define_enabled.ValidatableParameter, max_property_count.ValidatableParameter)
}
