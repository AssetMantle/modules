// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"github.com/AssetMantle/schema/go/properties"
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryRequest struct {
	PropertyList []properties.Property `json:"propertyList"`
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

func NewAuxiliaryRequest(propertyList ...properties.Property) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		PropertyList: propertyList,
	}
}
