// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintain

import (
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
)

type auxiliaryRequest struct {
	ids.ClassificationID `json:"classificationID" valid:"required~required field classificationID missing"`
	ids.IdentityID       `json:"identityID" valid:"required~required field identityID missing"`
	MaintainedProperties lists.PropertyList `json:"maintainedProperties" valid:"required~required field maintainedProperties missing"`
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

func NewAuxiliaryRequest(classificationID ids.ClassificationID, identityID ids.IdentityID, maintainedProperties lists.PropertyList) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ClassificationID:     classificationID,
		IdentityID:           identityID,
		MaintainedProperties: maintainedProperties,
	}
}
