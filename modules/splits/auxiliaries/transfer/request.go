// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
)

type auxiliaryRequest struct {
	FromID    ids.ID       `json:"fromID" valid:"required~required field fromID missing"`
	ToID      ids.ID       `json:"toID" valid:"required~required field toID missing"`
	OwnableID ids.ID       `json:"ownableID" valid:"required~required field ownableID missing"`
	Value     sdkTypes.Dec `json:"value" valid:"required~required field value missing"`
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

func NewAuxiliaryRequest(fromID ids.ID, toID ids.ID, ownableID ids.ID, value sdkTypes.Dec) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		FromID:    fromID,
		ToID:      toID,
		OwnableID: ownableID,
		Value:     value,
	}
}
