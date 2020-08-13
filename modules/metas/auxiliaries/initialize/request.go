/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package initialize

import (
	"github.com/asaskevich/govalidator"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryRequest struct {
	Data string `json:"data" valid:"required~required field data missing matches(^[A-Za-z]$)~invalid field data"`
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

func NewAuxiliaryRequest(data string) helpers.AuxiliaryRequest {
	return auxiliaryRequest{Data: data}
}
