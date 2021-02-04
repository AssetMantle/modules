/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

import (
	"github.com/asaskevich/govalidator"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type auxiliaryRequest struct {
	FromID               types.ID         `json:"fromID" valid:"required~required field fromID missing"`
	ToID                 types.ID         `json:"toID" valid:"required~required field toID missing"`
	ClassificationID     types.ID         `json:"classificationID" valid:"required~required field classificationID missing"`
	MaintainedProperties types.Properties `json:"maintainedProperties" valid:"required~required field maintainedProperties missing"`
	AddMaintainer        bool             `json:"addMaintainer" valid:"required~required field addMaintainer missing"`
	RemoveMaintainer     bool             `json:"removeMaintainer" valid:"required~required field removeMaintainer missing"`
	MutateMaintainer     bool             `json:"mutateMaintainer" valid:"required~required field mutateMaintainer missing"`
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(auxiliaryRequest)
	return Error
}

func auxiliaryRequestFromInterface(request helpers.AuxiliaryRequest) auxiliaryRequest {
	switch value := request.(type) {
	case auxiliaryRequest:
		return value
	default:
		return auxiliaryRequest{}
	}
}

func NewAuxiliaryRequest(fromID types.ID, toID types.ID, classificationID types.ID, maintainedProperties types.Properties, addMaintainer bool, removeMaintainer bool, mutateMaintainer bool) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		FromID:               fromID,
		ToID:                 toID,
		ClassificationID:     classificationID,
		MaintainedProperties: maintainedProperties,
		AddMaintainer:        addMaintainer,
		RemoveMaintainer:     removeMaintainer,
		MutateMaintainer:     mutateMaintainer,
	}
}
