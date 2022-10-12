// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
)

type auxiliaryRequest struct {
	FromID           ids.IdentityID       `json:"fromID" valid:"required~required field fromID missing"`
	ToID             ids.IdentityID       `json:"toID" valid:"required~required field toID missing"`
	ClassificationID ids.ClassificationID `json:"classificationID" valid:"required~required field classificationID missing"`
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

func NewAuxiliaryRequest(fromID ids.IdentityID, toID ids.IdentityID, classificationID ids.ClassificationID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		FromID:           fromID,
		ToID:             toID,
		ClassificationID: classificationID,
	}
}
