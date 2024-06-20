// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintain

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/qualified"
)

type auxiliaryRequest struct {
	MaintainedClassificationID ids.ClassificationID
	ids.IdentityID
	MaintainedMutables qualified.Mutables
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.MaintainedClassificationID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid maintained classification id %s", err.Error())
	}

	if err := auxiliaryRequest.IdentityID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid identity id %s", err.Error())
	}

	if err := auxiliaryRequest.MaintainedMutables.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid maintained mutables %s", err.Error())
	}

	return nil
}

func NewAuxiliaryRequest(maintainedClassificationID ids.ClassificationID, identityID ids.IdentityID, maintainedMutables qualified.Mutables) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		MaintainedClassificationID: maintainedClassificationID,
		IdentityID:                 identityID,
		MaintainedMutables:         maintainedMutables,
	}
}
