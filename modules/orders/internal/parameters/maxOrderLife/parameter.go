// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxOrderLife

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var ID = constantProperties.MaxOrderLifeProperty.GetKey()
var Parameter = baseParameters.NewParameter(baseProperties.NewMetaProperty(ID, baseData.NewHeightData(baseTypes.NewHeight(43210))))

func validator(i interface{}) error {
	switch value := i.(type) {
	case helpers.Parameter:
		if _, ok := value.GetMetaProperty().GetData().Get().(*baseData.HeightData); !ok || value.GetMetaProperty().GetID().GetKey().Compare(ID) != 0 {
			return errorConstants.IncorrectFormat
		}
		return nil
	case data.HeightData:
		if _, ok := i.(*baseData.HeightData); ok {
			return nil
		}
	}
	return errorConstants.IncorrectFormat
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
