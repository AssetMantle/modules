// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxProvisionAddressCount

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

var ID = constantProperties.MaxProvisionAddressCountProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewNumberData(16)))

func validator(i interface{}) error {
	var number *baseData.NumberData
	var ok bool
	switch value := i.(type) {
	case helpers.Parameter:
		number, ok = value.GetMetaProperty().GetData().Get().(*baseData.NumberData)
		if !ok || value.GetMetaProperty().GetID().GetKey().Compare(ID) != 0 {
			return errorConstants.IncorrectFormat.Wrapf("incorrect format for maxPropertyCount parameter, expected %T, got %T", baseData.NewNumberData(22), i)
		}
	case data.NumberData:
		number, ok = i.(*baseData.NumberData)
		if !ok {
			return errorConstants.IncorrectFormat.Wrapf("incorrect format for maxPropertyCount parameter, expected %T, got %T", baseData.NewNumberData(22), i)
		}
	default:
		return errorConstants.IncorrectFormat.Wrapf("incorrect format for maxPropertyCount parameter, expected %T, got %T", baseData.NewNumberData(22), i)
	}

	if number.Get() <= 0 {
		return errorConstants.IncorrectFormat.Wrapf("incorrect format for maxPropertyCount parameter, expected %T, got %T", baseData.NewNumberData(22), i)
	}

	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
