// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputizeAllowed

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

var ID = constantProperties.DeputizeAllowedProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewBooleanData(false)))

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

	return errorConstants.IncorrectFormat
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
