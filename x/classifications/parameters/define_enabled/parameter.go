// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define_enabled

import (
	baseData "github.com/AssetMantle/schema/go/data/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	"github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

var ID = constantProperties.DefineEnabledProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewBooleanData(false)))

func validator(i interface{}) error {
	switch value := i.(type) {
	case string:
		_, err := baseData.PrototypeBooleanData().FromString(value)
		return err
	default:
		return errorConstants.IncorrectFormat.Wrapf("incorrect type for defineEnabled parameter, expected %s type as string, got %T", baseData.NewBooleanData(false).GetTypeID().AsString(), i)
	}
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
