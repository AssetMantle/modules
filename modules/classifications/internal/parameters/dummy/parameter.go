// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
)

var ID = constantProperties.BondWeightProperty.GetKey()
var Parameter = baseTypes.NewParameter(base.NewMetaProperty(ID, baseData.NewDecData(sdkTypes.NewDec(2))))

func validator(i interface{}) error {
	switch value := i.(type) {
	case helpers.Parameter:
		if value.GetMetaProperty().GetID().Compare(ID) != 0 || value.GetMetaProperty().GetData().Get().(data.DecData).Get().IsNegative() {
			return errorConstants.InvalidParameter
		}

		return nil
	case data.DecData:
		if value.Get().IsNegative() {
			return errorConstants.InvalidParameter
		}

		return nil
	default:
		return errorConstants.IncorrectFormat
	}
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
