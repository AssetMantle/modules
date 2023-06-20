// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/constants"
	"github.com/AssetMantle/modules/x/identities/parameters/max_provision_address_count"
)

func Prototype() helpers.ParameterManager {
	return baseHelpers.NewParameterManager(constants.ModuleName, max_provision_address_count.ValidatableParameter)
}
