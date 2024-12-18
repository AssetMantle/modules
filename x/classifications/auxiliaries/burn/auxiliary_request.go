// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
)

type auxiliaryRequest struct {
	classificationID ids.ClassificationID
	bondAmount       math.Int
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.classificationID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid classification id: %s", err.Error())
	}

	if auxiliaryRequest.bondAmount.IsNegative() {
		return constants.InvalidRequest.Wrapf("bond amount cannot be negative")
	}

	return nil
}

func NewAuxiliaryRequest(classificationID ids.ClassificationID, bondAmount math.Int) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		classificationID: classificationID,
		bondAmount:       bondAmount,
	}
}
