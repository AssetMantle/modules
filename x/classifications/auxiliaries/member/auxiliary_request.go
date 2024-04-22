// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package member

import (
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/qualified"
)

type auxiliaryRequest struct {
	ids.ClassificationID
	qualified.Immutables
	qualified.Mutables
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.ClassificationID.ValidateBasic(); err != nil {
		return errorConstants.InvalidRequest.Wrapf("invalid classification id: %s", err.Error())
	}

	if err := auxiliaryRequest.Immutables.ValidateBasic(); err != nil {
		return errorConstants.InvalidRequest.Wrapf("invalid immutables: %s", err.Error())
	}

	if err := auxiliaryRequest.Mutables.ValidateBasic(); err != nil {
		return errorConstants.InvalidRequest.Wrapf("invalid mutables: %s", err.Error())
	}

	return nil
}

func NewAuxiliaryRequest(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ClassificationID: classificationID,
		Immutables:       immutables,
		Mutables:         mutables,
	}
}
