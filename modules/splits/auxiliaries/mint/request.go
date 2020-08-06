/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryRequest struct {
	OwnerID   types.ID     `json:"ownerID" valid:"required~required field ownerID missing"`
	OwnableID types.ID     `json:"ownableID" valid:"required~required field ownableID missing"`
	Split     sdkTypes.Dec `json:"split" valid:"required~required field assetID missing matches(^[0-9]$)~invalid field split"`
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func auxiliaryRequestFromInterface(AuxiliaryRequest helpers.AuxiliaryRequest) auxiliaryRequest {
	switch value := AuxiliaryRequest.(type) {
	case auxiliaryRequest:
		return value
	default:
		return auxiliaryRequest{}
	}
}

func NewAuxiliaryRequest(ownerID types.ID, ownableID types.ID, split sdkTypes.Dec) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		OwnerID:   base.NewID(ownerID.String()),
		OwnableID: base.NewID(ownableID.String()),
		Split:     split,
	}
}
