// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryRequest struct {
	OwnerID     ids.IdentityID `json:"ownerID" valid:"required~required field ownerID missing"`
	ids.AssetID `json:"assetID" valid:"required~required field assetId missing"`
	Value       sdkTypes.Int `json:"value" valid:"required~required field value missing"`
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

func NewAuxiliaryRequest(ownerID ids.IdentityID, assetID ids.AssetID, value sdkTypes.Int) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		OwnerID: ownerID,
		AssetID: assetID,
		Value:   value,
	}
}
