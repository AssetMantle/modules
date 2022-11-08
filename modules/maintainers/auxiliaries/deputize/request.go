// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
)

type auxiliaryRequest struct {
	FromID               ids.IdentityID `json:"fromID" valid:"required~required field fromID missing"`
	ToID                 ids.IdentityID `json:"toID" valid:"required~required field toID missing"`
	ids.ClassificationID `json:"classificationID" valid:"required~required field classificationID missing"`
	MaintainedProperties lists.PropertyList `json:"maintainedProperties" valid:"required~required field maintainedProperties missing"`
	CanMintAsset         bool               `json:"canMintAsset"`
	CanBurnAsset         bool               `json:"canBurnAsset"`
	CanRenumerateAsset   bool               `json:"canRenumerateAsset"`
	CanAddMaintainer     bool               `json:"canAddMaintainer"`
	CanRemoveMaintainer  bool               `json:"canRemoveMaintainer"`
	CanMutateMaintainer  bool               `json:"canMutateMaintainer"`
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

func NewAuxiliaryRequest(fromID ids.IdentityID, toID ids.IdentityID, classificationID ids.ClassificationID, maintainedProperties lists.PropertyList, canMintAsset bool, canBurnAsset bool, canRenumerateAsset bool, canAddMaintainer bool, canRemoveMaintainer bool, canMutateMaintainer bool) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		FromID:               fromID,
		ToID:                 toID,
		ClassificationID:     classificationID,
		MaintainedProperties: maintainedProperties,
		CanMintAsset:         canMintAsset,
		CanBurnAsset:         canBurnAsset,
		CanRenumerateAsset:   canRenumerateAsset,
		CanAddMaintainer:     canAddMaintainer,
		CanRemoveMaintainer:  canRemoveMaintainer,
		CanMutateMaintainer:  canMutateMaintainer,
	}
}
