// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authorize

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/go/ids"
)

type auxiliaryRequest struct {
	MaintainedClassificationID ids.ClassificationID
	MaintainerIdentityID       ids.IdentityID
	PermissionIDs              []ids.StringID
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.MaintainedClassificationID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid maintained classification id: %s ", err.Error())
	}

	if err := auxiliaryRequest.MaintainerIdentityID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid maintainer identity id: %s", err.Error())
	}

	for _, permissionID := range auxiliaryRequest.PermissionIDs {
		if err := permissionID.ValidateBasic(); err != nil {
			return constants.InvalidRequest.Wrapf("invalid permission id: %s", err.Error())
		}
	}

	return nil
}

func NewAuxiliaryRequest(maintainedClassificationID ids.ClassificationID, maintainerIdentityID ids.IdentityID, permissionIDs ...ids.StringID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		MaintainedClassificationID: maintainedClassificationID,
		MaintainerIdentityID:       maintainerIdentityID,
		PermissionIDs:              permissionIDs,
	}
}
