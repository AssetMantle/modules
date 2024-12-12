// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unwrap_allowed_coins

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/parameters"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	"github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var ID = constantProperties.UnwrapAllowedCoinsProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewListData(baseData.NewStringData(sdkTypes.DefaultBondDenom))))

func validator(parameter parameters.Parameter) error {
	if parameter.GetMetaProperty().GetID().Compare(Parameter.GetMetaProperty().GetID()) != 0 {
		return errorConstants.InvalidParameter.Wrapf("incorrect  ID, expected %s, got %s", ID.AsString(), parameter.GetMetaProperty().GetID().AsString())
	}

	if err := parameter.ValidateBasic(); err != nil {
		return errorConstants.InvalidParameter.Wrapf(err.Error())
	}

	listData := parameter.GetMetaProperty().GetData().Get().(*baseData.ListData)

	if err := listData.ValidateWithoutLengthCheck(); err != nil {
		return err
	}

	for _, anyData := range listData.Get() {
		if stringData, ok := anyData.Get().(*baseData.StringData); !ok {
			return errorConstants.IncorrectFormat.Wrapf("%s is not of type %s", anyData.Get().AsString(), baseData.PrototypeStringData().GetTypeID().AsString())
		} else if err := stringData.ValidateBasic(); err != nil {
			return err
		} else if err := sdkTypes.ValidateDenom(stringData.Get()); err != nil {
			return err
		}
	}
	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
