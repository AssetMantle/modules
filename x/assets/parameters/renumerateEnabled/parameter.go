// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerateEnabled

import (
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/parameters"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	"github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

var ID = constantProperties.RenumerateEnabledProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewBooleanData(false))) // NOTE: must always be set to false, legacy amino doesn't unmarshall false value well

func validator(i interface{}) error {
	switch value := i.(type) {
	case parameters.Parameter:
		if _, ok := value.GetMetaProperty().GetData().Get().(*baseData.BooleanData); ok && value.GetMetaProperty().GetID().GetKey().Compare(ID) == 0 {
			return nil
		}
	case data.BooleanData:
		if _, ok := i.(*baseData.BooleanData); ok {
			return nil
		}
	}

	return errorConstants.IncorrectFormat.Wrapf("incorrect format for renumerateEnabled parameter, expected %T, got %T", baseData.NewBooleanData(false), i)
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
