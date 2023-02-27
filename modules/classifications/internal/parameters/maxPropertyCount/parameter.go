// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxPropertyCount

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
)

var ID = constantProperties.MaxPropertyCountProperty.GetKey()
var Parameter = baseTypes.NewParameter(base.NewMetaProperty(ID, baseData.NewNumberData(22)))

func validator(i interface{}) error {
	if value, ok := i.(baseData.NumberData); ok && value.Get() > 0 {
		return nil
	}

	return errorConstants.IncorrectFormat.Wrapf("incorrect format for maxPropertyCount parameter, expected %T, got %T", baseData.NewNumberData(22), i)
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
