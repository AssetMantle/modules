// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/ids"
	"github.com/asaskevich/govalidator"
)

type auxiliaryRequest struct {
	MaintainedClassificationID ids.ClassificationID `json:"maintainedClassificationID" valid:"required~required field maintainedClassificationID missing"`
	MaintainerIdentityID       ids.IdentityID       `json:"maintainerIdentityID" valid:"required~required field maintainerIdentityID missing"`
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

func NewAuxiliaryRequest(maintainedClassificationID ids.ClassificationID, maintainedIdentityID ids.IdentityID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		MaintainedClassificationID: maintainedClassificationID,
		MaintainerIdentityID:       maintainedIdentityID,
	}
}
