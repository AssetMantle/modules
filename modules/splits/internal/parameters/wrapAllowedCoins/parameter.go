// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrapAllowedCoins

import (
	"github.com/AssetMantle/schema/x/data"
	baseData "github.com/AssetMantle/schema/x/data/base"
	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseParameters "github.com/AssetMantle/schema/x/parameters/base"
	"github.com/AssetMantle/schema/x/properties/base"
	constantProperties "github.com/AssetMantle/schema/x/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var ID = constantProperties.WrapAllowedCoinsProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewListData(baseData.NewIDData(baseIDs.NewCoinID(baseIDs.NewStringID(sdkTypes.DefaultBondDenom)).ToAnyID()))))

func validator(i interface{}) error {
	var listData *baseData.ListData
	var ok bool
	switch value := i.(type) {
	case helpers.Parameter:
		if listData, ok = value.GetMetaProperty().GetData().Get().(*baseData.ListData); !ok || value.GetMetaProperty().GetID().GetKey().Compare(ID) != 0 {
			return errorConstants.IncorrectFormat
		}
	case data.ListData:
		if listData, ok = i.(*baseData.ListData); !ok {
			return errorConstants.IncorrectFormat
		}
	default:
		return errorConstants.IncorrectFormat
	}

	for _, anyData := range listData.Get() {
		if idData, ok := anyData.Get().(*baseData.IDData); !ok {
			return errorConstants.IncorrectFormat
		} else if _, ok := idData.Get().Get().(ids.AnyOwnableID).Get().(*baseIDs.CoinID); !ok {
			return errorConstants.IncorrectFormat
		}
	}

	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
