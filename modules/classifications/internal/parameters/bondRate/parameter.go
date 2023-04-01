// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package bondRate

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
)

var ID = constantProperties.BondRateProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewNumberData(1)))

func validator(i interface{}) error {
	var number *baseData.NumberData
	var ok bool
	switch value := i.(type) {
	case helpers.Parameter:
		number, ok = value.GetMetaProperty().GetData().Get().(*baseData.NumberData)
		if !ok || value.GetMetaProperty().GetID().GetKey().Compare(ID) != 0 {
			return errorConstants.IncorrectFormat
		}
	case data.NumberData:
		number, ok = i.(*baseData.NumberData)
		if !ok {
			return errorConstants.IncorrectFormat
		}
	default:
		return errorConstants.IncorrectFormat
	}

	if number.Get() <= 0 {
		return errorConstants.IncorrectFormat
	}

	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
