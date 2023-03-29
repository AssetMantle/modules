// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/qualified"
)

type auxiliaryRequest struct {
	sdkTypes.AccAddress
	qualified.Immutables `json:"immutables" valid:"required~required field immutableProperties missing"`
	qualified.Mutables   `json:"mutables" valid:"required~required field mutableProperties missing"`
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

func NewAuxiliaryRequest(accAddress sdkTypes.AccAddress, immutables qualified.Immutables, mutables qualified.Mutables) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		AccAddress: accAddress,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
