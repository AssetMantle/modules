// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"
)

type auxiliaryRequest struct {
	FromIdentityID             ids.IdentityID
	ToIdentityID               ids.IdentityID
	MaintainedClassificationID ids.ClassificationID
	MaintainedProperties       lists.PropertyList
	CanAddMaintainer           bool
	CanRemoveMaintainer        bool
	CanMutateMaintainer        bool
	PermissionIDs              []ids.StringID
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.FromIdentityID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid from identity id: %s", err.Error())
	}

	if err := auxiliaryRequest.ToIdentityID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid to identity id: %s", err.Error())
	}

	if err := auxiliaryRequest.MaintainedClassificationID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid maintained classification id: %s", err.Error())
	}

	if err := auxiliaryRequest.MaintainedProperties.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid maintained properties: %s", err.Error())
	}

	for _, permissionID := range auxiliaryRequest.PermissionIDs {
		if err := permissionID.ValidateBasic(); err != nil {
			return constants.InvalidRequest.Wrapf("invalid permission id: %s", err.Error())
		}
	}

	return nil
}

func NewAuxiliaryRequest(fromIdentityID ids.IdentityID, toIdentityID ids.IdentityID, maintainedClassificationID ids.ClassificationID, maintainedProperties lists.PropertyList, canAddMaintainer bool, canRemoveMaintainer bool, canMutateMaintainer bool, permissionIDs ...ids.StringID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		FromIdentityID:             fromIdentityID,
		ToIdentityID:               toIdentityID,
		MaintainedClassificationID: maintainedClassificationID,
		MaintainedProperties:       maintainedProperties,
		CanAddMaintainer:           canAddMaintainer,
		CanRemoveMaintainer:        canRemoveMaintainer,
		CanMutateMaintainer:        canMutateMaintainer,
		PermissionIDs:              permissionIDs,
	}
}
