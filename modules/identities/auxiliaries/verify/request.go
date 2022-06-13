// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
)

type auxiliaryRequest struct {
	Address    sdkTypes.AccAddress `json:"address" valid:"required~required field address missing, matches(^[a-z0-9]*$)~field address is invalid"`
	IdentityID ids.ID              `json:"identityID" valid:"required~required field identityID missing"`
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

func NewAuxiliaryRequest(address sdkTypes.AccAddress, identityID ids.ID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		Address:    address,
		IdentityID: identityID,
	}
}
