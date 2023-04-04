// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxProvisionAddressCount

import (
	"github.com/AssetMantle/schema/x/data"
	baseData "github.com/AssetMantle/schema/x/data/base"
	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	baseParameters "github.com/AssetMantle/schema/x/parameters/base"
	"github.com/AssetMantle/schema/x/properties/base"
	constantProperties "github.com/AssetMantle/schema/x/properties/constants"
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
