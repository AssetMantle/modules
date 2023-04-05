// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxOrderLife

import (
	"github.com/AssetMantle/schema/x/data"
	baseData "github.com/AssetMantle/schema/x/data/base"
	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/AssetMantle/schema/x/parameters"
	baseParameters "github.com/AssetMantle/schema/x/parameters/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	constantProperties "github.com/AssetMantle/schema/x/properties/constants"
	baseTypes "github.com/AssetMantle/schema/x/types/base"
)

var ID = constantProperties.MaxOrderLifeProperty.GetKey()
var Parameter = baseParameters.NewParameter(baseProperties.NewMetaProperty(ID, baseData.NewHeightData(baseTypes.NewHeight(43210))))

func validator(i interface{}) error {
	var height *baseData.HeightData
	var ok bool
	switch value := i.(type) {
	case parameters.Parameter:
		height, ok = value.GetMetaProperty().GetData().Get().(*baseData.HeightData)
		if !ok || value.GetMetaProperty().GetID().GetKey().Compare(ID) != 0 {
			return errorConstants.IncorrectFormat
		}
	case data.HeightData:
		height, ok = i.(*baseData.HeightData)
		if !ok {
			return errorConstants.IncorrectFormat
		}
	default:
		return errorConstants.IncorrectFormat
	}

	if height.Get().Get() < -1 {
		return errorConstants.IncorrectFormat
	}

	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
