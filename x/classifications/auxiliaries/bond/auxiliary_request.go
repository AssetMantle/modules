// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package bond

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryRequest struct {
	classificationID ids.ClassificationID
	accAddress       sdkTypes.AccAddress
	bondAmount       sdkTypes.Int
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

func NewAuxiliaryRequest(classificationID ids.ClassificationID, fromAddress sdkTypes.AccAddress, bondAmount sdkTypes.Int) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		classificationID: classificationID,
		accAddress:       fromAddress,
		bondAmount:       bondAmount,
	}
}
