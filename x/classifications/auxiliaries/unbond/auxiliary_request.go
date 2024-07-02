// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unbond

import (
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryRequest struct {
	classificationID ids.ClassificationID
	accAddress       sdkTypes.AccAddress
	bondAmount       sdkTypes.Int
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.classificationID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid classification id: %s", err.Error())
	}

	if auxiliaryRequest.accAddress.Empty() {
		return constants.InvalidRequest.Wrapf("address cannot be empty")
	}

	if auxiliaryRequest.bondAmount.IsNegative() {
		return constants.InvalidRequest.Wrapf("bond amount cannot be negative")
	}

	return nil
}

func NewAuxiliaryRequest(classificationID ids.ClassificationID, fromAddress sdkTypes.AccAddress, bondAmount sdkTypes.Int) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		classificationID: classificationID,
		accAddress:       fromAddress,
		bondAmount:       bondAmount,
	}
}
