// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package max_order_life

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseData "github.com/AssetMantle/schema/data/base"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

var ID = constantProperties.MaxOrderLifeProperty.GetKey()
var Parameter = baseParameters.NewParameter(baseProperties.NewMetaProperty(ID, baseData.NewHeightData(baseTypes.NewHeight(43210))))

func validator(i interface{}) error {
	switch value := i.(type) {
	case string:
		if height, err := baseData.PrototypeHeightData().FromString(value); err != nil {
			return err
		} else {
			err = height.(*baseData.HeightData).ValidateBasic()
			return err
		}
	default:
		return errorConstants.IncorrectFormat.Wrapf("incorrect type for maxOrderLife parameter, expected %s type as string, got %T", baseData.NewHeightData(baseTypes.NewHeight(0)).GetTypeID().AsString(), i)
	}
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
