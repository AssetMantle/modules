/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"github.com/asaskevich/govalidator"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type auxiliaryRequest struct {
	ImmutableTraits types.Immutables `json:"immutableTraits" valid:"required~required field immutableTraits missing"`
	MutableTraits   types.Mutables   `json:"mutableTraits" valid:"required~required field mutableTraits missing"`
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

func NewAuxiliaryRequest(immutableTraits types.Immutables, mutableTraits types.Mutables) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ImmutableTraits: immutableTraits,
		MutableTraits:   mutableTraits,
	}
}
