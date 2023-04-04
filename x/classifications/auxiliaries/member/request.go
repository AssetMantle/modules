// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package member

import (
	"github.com/asaskevich/govalidator"

	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/ids"
	"github.com/AssetMantle/schema/x/qualified"
)

type auxiliaryRequest struct {
	ids.ClassificationID `json:"classificationID" valid:"required~required field classificationID missing"`
	qualified.Immutables `json:"immutableProperties"`
	qualified.Mutables   `json:"mutableProperties"`
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
		panic(errorConstants.InvalidRequest.Wrapf("invalid request type %T", value))
	}
}

func NewAuxiliaryRequest(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ClassificationID: classificationID,
		Immutables:       immutables,
		Mutables:         mutables,
	}
}
