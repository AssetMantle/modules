// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrapAllowedCoins

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var ID = constantProperties.WrapAllowedCoinsProperty.GetKey()
var Parameter = baseTypes.NewParameter(base.NewMetaProperty(ID, baseData.NewListData(baseData.NewIDData(baseIDs.NewCoinID(baseIDs.NewStringID(sdkTypes.DefaultBondDenom))))))

func validData(listData *baseData.ListData) bool {
	for _, anyData := range listData.Get() {
		if idData, ok := anyData.Get().(*baseData.IDData); !ok {
			return false
		} else if _, ok := idData.Get().Get().(ids.AnyOwnableID).Get().(*baseIDs.CoinID); !ok {
			return false
		}
	}
	return true
}

func validator(i interface{}) error {
	switch value := i.(type) {
	case helpers.Parameter:
		if listData, ok := value.GetMetaProperty().GetData().Get().(*baseData.ListData); !ok || value.GetMetaProperty().GetID().GetKey().Compare(ID) != 0 {
			return errorConstants.IncorrectFormat
		} else if !validData(listData) {
			return errorConstants.IncorrectFormat
		}
	case data.ListData:
		if listData, ok := i.(*baseData.ListData); !ok {
			return errorConstants.IncorrectFormat
		} else if !validData(listData) {
			return errorConstants.IncorrectFormat
		}
	default:
		return errorConstants.IncorrectFormat
	}

	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
