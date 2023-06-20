// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap_allowed_coins

import (
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	"github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

var ID = constantProperties.WrapAllowedCoinsProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewListData(baseData.NewIDData(baseIDs.NewCoinID(baseIDs.NewStringID(sdkTypes.DefaultBondDenom)).ToAnyID()))))

func validator(i interface{}) error {
	switch value := i.(type) {
	case string:
		if data, err := baseData.PrototypeListData().FromString(value); err != nil {
			return err
		} else if err = ValidateWrapAllowedCoinsProperty(data.(*baseData.ListData)); err != nil {
			return err
		} else {
			return data.(*baseData.ListData).ValidateBasic()
		}
	default:
		return errorConstants.IncorrectFormat
	}
}

func ValidateWrapAllowedCoinsProperty(listData data.ListData) error {
	for _, anyData := range listData.Get() {
		if idData, ok := anyData.Get().(*baseData.IDData); !ok {
			return errorConstants.IncorrectFormat
		} else if coinID, ok := idData.Get().Get().(ids.AnyOwnableID).Get().(*baseIDs.CoinID); !ok {
			return errorConstants.IncorrectFormat
		} else if err := coinID.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
