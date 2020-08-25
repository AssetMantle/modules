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
	ImmutableMetaTraits types.MetaProperties `json:"immutableMetaTraits" valid:"required~required field immutableMetaTraits missing"`
	ImmutableTraits     types.Properties     `json:"immutableTraits" valid:"required~required field immutableTraits missing"`
	MutableMetaTraits   types.MetaProperties `json:"mutableMetaTraits" valid:"required~required field mutableMetaTraits missing"`
	MutableTraits       types.Properties     `json:"mutableTraits" valid:"required~required field mutableTraits missing"`
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

func NewAuxiliaryRequest(immutableMetaTraits types.MetaProperties, immutableTraits types.Properties, mutableMetaTraits types.MetaProperties, mutableTraits types.Properties) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ImmutableMetaTraits: immutableMetaTraits,
		ImmutableTraits:     immutableTraits,
		MutableMetaTraits:   mutableMetaTraits,
		MutableTraits:       mutableTraits,
	}
}
