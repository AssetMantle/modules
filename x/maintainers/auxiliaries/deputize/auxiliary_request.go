// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/lists"
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryRequest struct {
	FromIdentityID             ids.IdentityID       `json:"fromIdentityID" valid:"required~required field fromIdentityID missing"`
	ToIdentityID               ids.IdentityID       `json:"toIdentityID" valid:"required~required field toIdentityID missing"`
	MaintainedClassificationID ids.ClassificationID `json:"maintainedClassificationID" valid:"required~required field maintainedClassificationID missing"`
	MaintainedProperties       lists.PropertyList   `json:"maintainedProperties" valid:"required~required field maintainedProperties missing"`
	CanAddMaintainer           bool                 `json:"canAddMaintainer" valid:"required~required field canAddMaintainer missing"`
	CanRemoveMaintainer        bool                 `json:"canRemoveMaintainer" valid:"required~required field canRemoveMaintainer missing"`
	CanMutateMaintainer        bool                 `json:"canMutateMaintainer" valid:"required~required field canMutateMaintainer missing"`
	PermissionIDs              []ids.StringID       `json:"permissionIDs" valid:"required~required field permissionIDs missing"`
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(auxiliaryRequest)
	return err
}

func auxiliaryRequestFromInterface(request helpers.AuxiliaryRequest) auxiliaryRequest {
	switch value := request.(type) {
	case auxiliaryRequest:
		return value
	default:
		return auxiliaryRequest{}
	}
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
