// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
)

type auxiliaryRequest struct {
	OwnerID   ids.IdentityID `json:"ownerID" valid:"required~required field ownerID missing"`
	OwnableID ids.OwnableID  `json:"ownableID" valid:"required~required field ownableID missing"`
	Value     sdkTypes.Dec   `json:"value" valid:"required~required field value missing"`
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

func NewAuxiliaryRequest(ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		OwnerID:   ownerID,
		OwnableID: ownableID,
		Value:     value,
	}
}
