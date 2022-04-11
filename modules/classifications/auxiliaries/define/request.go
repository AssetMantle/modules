// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/lists"
)

type auxiliaryRequest struct {
	ImmutableProperties lists.PropertyList `json:"immutableProperties" valid:"required~required field immutableProperties missing"`
	MutableProperties   lists.PropertyList `json:"mutableProperties" valid:"required~required field mutableProperties missing"`
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

func NewAuxiliaryRequest(immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ImmutableProperties: immutableProperties,
		MutableProperties:   mutableProperties,
	}
}
