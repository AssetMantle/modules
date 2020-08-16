/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package conform

import (
	"github.com/asaskevich/govalidator"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type auxiliaryRequest struct {
	ClassificationID types.ID         `json:"classificationID" valid:"required~required field classificationID missing"`
	Mutables         types.Mutables   `json:"mutables"`
	Immutables       types.Immutables `json:"immutables"`
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(auxiliaryRequest)
	return Error
}

func auxiliaryRequestFromInterface(AuxiliaryRequest helpers.AuxiliaryRequest) auxiliaryRequest {
	switch value := AuxiliaryRequest.(type) {
	case auxiliaryRequest:
		return value
	default:
		return auxiliaryRequest{}
	}
}

func NewAuxiliaryRequest(classificationID types.ID, mutables types.Mutables, immutables types.Immutables) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ClassificationID: classificationID,
		Mutables:         mutables,
		Immutables:       immutables,
	}
}
