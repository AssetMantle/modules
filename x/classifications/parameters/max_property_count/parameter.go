// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package max_property_count

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/parameters"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	"github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	"github.com/cosmos/cosmos-sdk/types"
)

var ID = constantProperties.MaxPropertyCountProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewNumberData(types.NewInt(22))))

func validator(parameter parameters.Parameter) error {
	if parameter.GetMetaProperty().GetID().Compare(Parameter.GetMetaProperty().GetID()) != 0 {
		return errorConstants.InvalidParameter.Wrapf("incorrect  ID, expected %s, got %s", ID.AsString(), parameter.GetMetaProperty().GetID().AsString())
	}

	if err := parameter.ValidateBasic(); err != nil {
		return errorConstants.InvalidParameter.Wrapf(err.Error())
	}

	if parameter.GetMetaProperty().GetData().Get().(*baseData.NumberData).Get().LT(types.OneInt()) {
		return errorConstants.InvalidParameter.Wrapf("%s must be greater than or equal to 1", ID.AsString())
	}

	if parameter.GetMetaProperty().GetData().Get().(*baseData.NumberData).Get().GT(types.NewInt(44)) {
		return errorConstants.InvalidParameter.Wrapf("%s must be less than than or equal to 44", ID.AsString())
	}

	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
