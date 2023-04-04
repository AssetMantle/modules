// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/x/maintainers/internal/module"
	"github.com/AssetMantle/modules/x/maintainers/internal/parameters/deputizeAllowed"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(module.Name, deputizeAllowed.ValidatableParameter)
}
