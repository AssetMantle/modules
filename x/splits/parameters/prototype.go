// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/parameters/transfer_enabled"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(transfer_enabled.ValidatableParameter)
}
