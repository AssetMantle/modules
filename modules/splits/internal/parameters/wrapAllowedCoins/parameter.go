// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrapAllowedCoins

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
)

var ID = constantProperties.WrapAllowedCoinsProperty.GetKey()
var Parameter = baseTypes.NewParameter(base.NewMetaProperty(ID, baseData.NewListData(baseData.NewIDData(baseIDs.NewCoinID(baseIDs.NewStringID(sdkTypes.DefaultBondDenom))))))

func validator(i interface{}) error {
	if value, ok := i.(baseData.ListData); ok {
		return errorConstants.IncorrectFormat
	} else {
		for _, anyData := range value.Get() {
			if idData, ok := anyData.Get().(*baseData.IDData); !ok {
				return errorConstants.IncorrectFormat
			} else if _, ok := idData.Get().Get().(*baseIDs.CoinID); !ok {
				return errorConstants.IncorrectFormat
			}
		}
	}
	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
