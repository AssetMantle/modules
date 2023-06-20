// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type auxiliaryRequest struct {
	FromID    ids.IdentityID `json:"fromID" valid:"required~required field fromID missing"`
	ToID      ids.IdentityID `json:"toID" valid:"required~required field toID missing"`
	OwnableID ids.OwnableID  `json:"ownableID" valid:"required~required field ownableID missing"`
	Value     sdkTypes.Int   `json:"value" valid:"required~required field value missing"`
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

func NewAuxiliaryRequest(fromID ids.IdentityID, toID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Int) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		FromID:    fromID,
		ToID:      toID,
		OwnableID: ownableID,
		Value:     value,
	}
}
