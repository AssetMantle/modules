// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintain

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/qualified"
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryRequest struct {
	MaintainedClassificationID ids.ClassificationID `json:"maintainedClassificationID" valid:"required~required field maintainedClassificationID missing"`
	ids.IdentityID             `json:"identityID" valid:"required~required field identityID missing"`
	MaintainedMutables         qualified.Mutables `json:"maintainedMutables" valid:"required~required field maintainedProperties missing"`
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

func NewAuxiliaryRequest(maintainedClassificationID ids.ClassificationID, identityID ids.IdentityID, maintainedMutables qualified.Mutables) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		MaintainedClassificationID: maintainedClassificationID,
		IdentityID:                 identityID,
		MaintainedMutables:         maintainedMutables,
	}
}
