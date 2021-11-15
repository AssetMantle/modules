/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package verify

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type auxiliaryRequest struct {
	Address    sdkTypes.AccAddress `json:"address" valid:"required~required field address missing matches(^mantle[a-z0-9]{39}$)~field address is invalid"`
	IdentityID types.ID            `json:"identityID" valid:"required~required field identityID missing"`
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

func NewAuxiliaryRequest(address sdkTypes.AccAddress, identityID types.ID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		Address:    address,
		IdentityID: identityID,
	}
}
