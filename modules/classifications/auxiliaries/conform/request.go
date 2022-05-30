// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package conform

import (
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
)

type auxiliaryRequest struct {
	ClassificationID    ids.ID             `json:"classificationID" valid:"required~required field classificationID missing"`
	ImmutableProperties lists.PropertyList `json:"immutableProperties"`
	MutableProperties   lists.PropertyList `json:"mutableProperties"`
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

func NewAuxiliaryRequest(classificationID ids.ID, immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ClassificationID:    classificationID,
		ImmutableProperties: immutableProperties,
		MutableProperties:   mutableProperties,
	}
}
