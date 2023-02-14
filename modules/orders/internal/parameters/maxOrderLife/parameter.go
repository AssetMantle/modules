// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxOrderLife

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var ID = constantProperties.MaxOrderLifeProperty.GetKey()
var Parameter = baseParameters.NewParameter(baseProperties.NewMetaProperty(ID, baseData.NewHeightData(baseTypes.NewHeight(43210))))

func validator(i interface{}) error {
	if value, ok := i.(baseData.HeightData); ok && value.Get().Get() > -1 {
		return nil
	}

	return errorConstants.IncorrectFormat
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
