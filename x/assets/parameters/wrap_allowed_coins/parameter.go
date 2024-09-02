// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap_allowed_coins

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	"github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var ID = constantProperties.WrapAllowedCoinsProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewListData(baseData.NewStringData(sdkTypes.DefaultBondDenom))))

func validator(i interface{}) error {
	switch value := i.(type) {
	case string:
		if data, err := baseData.PrototypeListData().FromString(value); err != nil {
			return err
		} else if err = validateWrapAllowedCoinsProperty(data.ToAnyData().Get().(*baseData.ListData)); err != nil {
			return err
		} else {
			return data.(*baseData.ListData).ValidateWithoutLengthCheck()
		}
	default:
		return errorConstants.IncorrectFormat.Wrapf("incorrect type for wrapAllowedCoins parameter, expected %s type as string, got %T", baseData.PrototypeListData().GetTypeID().AsString(), i)
	}
}

func validateWrapAllowedCoinsProperty(listData data.ListData) error {
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
