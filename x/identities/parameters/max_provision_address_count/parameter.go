// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package max_provision_address_count

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/data/constants"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	"github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var ID = constantProperties.MaxProvisionAddressCountProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewNumberData(sdkTypes.NewInt(16))))

func validator(i interface{}) error {
	switch value := i.(type) {
	case string:
		if number, err := baseData.PrototypeNumberData().FromString(value); err != nil {
			return err
		} else if number.(*baseData.NumberData).Get().Equal(sdkTypes.ZeroInt()) {
			return errorConstants.IncorrectFormat.Wrapf("maxProvisionAddressCount parameter cannot be zero")
		} else if number.(*baseData.NumberData).Get().GT(sdkTypes.NewInt(constants.MaxListLength)) {
			return errorConstants.IncorrectFormat.Wrapf("maxProvisionAddressCount parameter cannot be greater than %d", constants.MaxListLength)
		} else {
			err = number.(*baseData.NumberData).ValidateBasic()
			return err
		}
	default:
		return errorConstants.IncorrectFormat.Wrapf("incorrect type for maxProvisionAddressCount parameter, expected %s type as string, got %T", baseData.NewNumberData(sdkTypes.OneInt()).GetTypeID().AsString(), i)
	}
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
