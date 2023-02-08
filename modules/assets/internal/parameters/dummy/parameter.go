// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties/base"
)

var ID = baseIDs.NewStringID("dummy")
var Parameter = baseTypes.NewParameter(base.NewMetaProperty(ID, baseData.NewDecData(sdkTypes.ZeroDec())))

func validator(i interface{}) error {
	if value, ok := i.(baseData.DecData); ok && !value.Get().IsNegative() {
		return nil
	}

	return constants.IncorrectFormat

}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
