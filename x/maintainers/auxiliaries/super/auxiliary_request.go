// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/qualified"
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryRequest struct {
	MaintainedClassificationID ids.ClassificationID `json:"maintainedClassificationID" valid:"required~required field maintainedClassificationID missing"`
	ToIdentityID               ids.IdentityID       `json:"toIdentityID" valid:"required~required field identityID missing"`
	MaintainedMutables         qualified.Mutables   `json:"maintainedMutables" valid:"required~required field maintainedMutables missing"`
	PermissionIDs              []ids.StringID       `json:"permissionIDs"`
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

func NewAuxiliaryRequest(maintainedClassificationID ids.ClassificationID, toIdentityID ids.IdentityID, maintainedMutables qualified.Mutables, permissionIDs ...ids.StringID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		MaintainedClassificationID: maintainedClassificationID,
		ToIdentityID:               toIdentityID,
		MaintainedMutables:         maintainedMutables,
		PermissionIDs:              permissionIDs,
	}
}
