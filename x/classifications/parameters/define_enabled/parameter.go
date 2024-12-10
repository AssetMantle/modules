// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define_enabled

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/parameters"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	"github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
)

var ID = constantProperties.DefineEnabledProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewBooleanData(true)))

func validator(parameter parameters.Parameter) error {
	if parameter.GetMetaProperty().GetID().Compare(Parameter.GetMetaProperty().GetID()) != 0 {
		return errorConstants.InvalidParameter.Wrapf("incorrect  ID, expected %s, got %s", ID.AsString(), parameter.GetMetaProperty().GetID().AsString())
	}

	if err := parameter.ValidateBasic(); err != nil {
		return errorConstants.InvalidParameter.Wrapf(err.Error())
	}

	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
