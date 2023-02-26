// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burnEnabled

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
)

var ID = constantProperties.BurnEnabledProperty.GetKey()
var Parameter = baseTypes.NewParameter(base.NewMetaProperty(ID, baseData.NewBooleanData(false))) // NOTE: must always be set to false, legacy amino doesn't unmarshall false value well

func validator(i interface{}) error {
	if _, ok := i.(baseData.BooleanData); ok {
		return nil
	}

	return errorConstants.IncorrectFormat

}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
