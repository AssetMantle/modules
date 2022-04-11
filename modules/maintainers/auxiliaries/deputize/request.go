// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
)

type auxiliaryRequest struct {
	FromID               types.ID           `json:"fromID" valid:"required~required field fromID missing"`
	ToID                 types.ID           `json:"toID" valid:"required~required field toID missing"`
	ClassificationID     types.ID           `json:"classificationID" valid:"required~required field classificationID missing"`
	MaintainedProperties lists.PropertyList `json:"maintainedProperties" valid:"required~required field maintainedProperties missing"`
	AddMaintainer        bool               `json:"addMaintainer"`
	RemoveMaintainer     bool               `json:"removeMaintainer"`
	MutateMaintainer     bool               `json:"mutateMaintainer"`
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

func NewAuxiliaryRequest(fromID types.ID, toID types.ID, classificationID types.ID, maintainedProperties lists.PropertyList, addMaintainer bool, removeMaintainer bool, mutateMaintainer bool) helpers.AuxiliaryRequest {
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
