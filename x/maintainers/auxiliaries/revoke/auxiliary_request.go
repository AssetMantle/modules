// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/go/ids"
)

type auxiliaryRequest struct {
	FromID                     ids.IdentityID
	ToID                       ids.IdentityID
	MaintainedClassificationID ids.ClassificationID
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.FromID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid from id: %s", err.Error())
	}

	if err := auxiliaryRequest.ToID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid to id: %s", err.Error())
	}

	if err := auxiliaryRequest.MaintainedClassificationID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid maintained classification id: %s", err.Error())
	}

	return nil
}

func NewAuxiliaryRequest(fromID ids.IdentityID, toID ids.IdentityID, maintainedClassificationID ids.ClassificationID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		FromID:                     fromID,
		ToID:                       toID,
		MaintainedClassificationID: maintainedClassificationID,
	}
}
