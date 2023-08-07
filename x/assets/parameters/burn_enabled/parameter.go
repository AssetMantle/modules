// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn_enabled

import (
	baseData "github.com/AssetMantle/schema/go/data/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	"github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

var ID = constantProperties.BurnEnabledProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewBooleanData(true))) // NOTE: must always be set to false, legacy amino doesn't unmarshall false value well

func validator(i interface{}) error {
	switch value := i.(type) {
	case string:
		_, err := baseData.PrototypeBooleanData().FromString(value)
		return err
	default:
		return errorConstants.IncorrectFormat.Wrapf("incorrect format for burnEnabled parameter, expected %T, got %T", baseData.NewBooleanData(false), i)
	}
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
