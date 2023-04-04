// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/modules/metas/internal/module"
	"github.com/AssetMantle/modules/modules/metas/internal/parameters/revealEnabled"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(module.Name, revealEnabled.ValidatableParameter)
}
