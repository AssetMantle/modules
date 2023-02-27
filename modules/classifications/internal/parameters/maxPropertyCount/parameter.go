// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxPropertyCount

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
)

var ID = constantProperties.MaxPropertyCountProperty.GetKey()
var Parameter = baseTypes.NewParameter(base.NewMetaProperty(ID, baseData.NewNumberData(22)))

func validator(i interface{}) error {
	switch value := i.(type) {
	case helpers.Parameter:
		if _, ok := value.GetMetaProperty().GetData().Get().(*baseData.NumberData); !ok || value.GetMetaProperty().GetID().GetKey().Compare(ID) != 0 {
			return errorConstants.IncorrectFormat
		}
		return nil
	case data.NumberData:
		if _, ok := i.(*baseData.NumberData); ok {
			return nil
		}
	}

	return errorConstants.IncorrectFormat
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
