// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type auxiliaryRequest struct {
	OwnerID   ids.ID       `json:"ownerID" valid:"required~required field ownerID missing"`
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

func NewAuxiliaryRequest(ownerID fmt.Stringer, ownableID fmt.Stringer, value sdkTypes.Dec) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		OwnerID:   baseIDs.NewStringID(ownerID.String()),
		OwnableID: baseIDs.NewStringID(ownableID.String()),
		Value:     value,
	}
}
