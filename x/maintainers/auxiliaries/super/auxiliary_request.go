// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/qualified"
)

type auxiliaryRequest struct {
	MaintainedClassificationID ids.ClassificationID
	ToIdentityID               ids.IdentityID
	MaintainedMutables         qualified.Mutables
	PermissionIDs              []ids.StringID
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.MaintainedClassificationID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid maintained classification id: %s", err.Error())
	}

	if err := auxiliaryRequest.ToIdentityID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid to identity id: %s", err.Error())
	}

	if err := auxiliaryRequest.MaintainedMutables.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid maintained mutables: %s", err.Error())
	}

	for _, permissionID := range auxiliaryRequest.PermissionIDs {
		if err := permissionID.ValidateBasic(); err != nil {
			return constants.InvalidRequest.Wrapf("invalid permission id: %s", err.Error())
		}
	}

	return nil

}

func NewAuxiliaryRequest(maintainedClassificationID ids.ClassificationID, toIdentityID ids.IdentityID, maintainedMutables qualified.Mutables, permissionIDs ...ids.StringID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		MaintainedClassificationID: maintainedClassificationID,
		ToIdentityID:               toIdentityID,
		MaintainedMutables:         maintainedMutables,
		PermissionIDs:              permissionIDs,
	}
}
