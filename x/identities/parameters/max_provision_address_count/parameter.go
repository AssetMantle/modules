// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package max_provision_address_count

import (
	"cosmossdk.io/math"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseData "github.com/AssetMantle/schema/data/base"
	dataConstants "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/parameters"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	"github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var ID = constantProperties.MaxProvisionAddressCountProperty.GetKey()
var Parameter = baseParameters.NewParameter(base.NewMetaProperty(ID, baseData.NewNumberData(sdkTypes.NewInt(16))))

func validator(parameter parameters.Parameter) error {
	if parameter.GetMetaProperty().GetID().Compare(Parameter.GetMetaProperty().GetID()) != 0 {
		return errorConstants.InvalidParameter.Wrapf("incorrect  ID, expected %s, got %s", ID.AsString(), parameter.GetMetaProperty().GetID().AsString())
	}

	if err := parameter.ValidateBasic(); err != nil {
		return errorConstants.InvalidParameter.Wrapf(err.Error())
	}

	if parameter.GetMetaProperty().GetData().Get().(*baseData.NumberData).Get().LT(sdkTypes.OneInt()) {
		return errorConstants.InvalidParameter.Wrapf("%s must be greater than or equal to 1", ID.AsString())
	}

	if parameter.GetMetaProperty().GetData().Get().(*baseData.NumberData).Get().GT(math.NewInt(dataConstants.MaxListLength)) {
		return errorConstants.InvalidParameter.Wrapf("%s must be less than or equal to %d", ID.AsString(), dataConstants.MaxListLength)
	}

	return nil
}

var ValidatableParameter = baseHelpers.NewValidatableParameter(Parameter, validator)
