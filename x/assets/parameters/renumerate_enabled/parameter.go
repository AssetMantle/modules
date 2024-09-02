// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate_enabled

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseData "github.com/AssetMantle/schema/data/base"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	"github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
)

var ID = constantProperties.RenumerateEnabledProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewBooleanData(true)))

func validator(i interface{}) error {
	switch value := i.(type) {
	case string:
		_, err := baseData.PrototypeBooleanData().FromString(value)
		return err
	default:
		return errorConstants.IncorrectFormat.Wrapf("incorrect type for renumerateEnabled parameter, expected %s type as string, got %T", baseData.NewBooleanData(false).GetTypeID().AsString(), i)
	}
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
