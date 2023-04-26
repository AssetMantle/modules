// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxProvisionAddressCount

import (
	baseData "github.com/AssetMantle/schema/go/data/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	"github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

var ID = constantProperties.MaxProvisionAddressCountProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewNumberData(16)))

func validator(i interface{}) error {
	switch value := i.(type) {
	case string:
		if number, err := baseData.PrototypeNumberData().FromString(value); err != nil {
			return err
		} else if number.(*baseData.NumberData).Get() <= 0 {
			return errorConstants.IncorrectFormat.Wrapf("incorrect format for maxProvisionAddressCount parameter, has to be a positive whole number")
		} else {
			err = number.(*baseData.NumberData).ValidateBasic()
			return err
		}
	default:
		return errorConstants.IncorrectFormat.Wrapf("incorrect format for maxPropertyCount parameter, expected %T, got %T", baseData.NewNumberData(22), i)
	}

	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
