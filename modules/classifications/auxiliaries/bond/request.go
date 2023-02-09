// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package bond

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/AssetMantle/modules/schema/helpers"
)

type auxiliaryRequest struct {
	ClassificationID ids.ClassificationID
	address          sdkTypes.AccAddress
	moduleName       string
	bankKeeper       bankKeeper.Keeper
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

func NewAuxiliaryRequest(classificationID ids.ClassificationID, fromAddress sdkTypes.AccAddress, moduleName string, bankKeeper bankKeeper.Keeper) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ClassificationID: classificationID,
		address:          fromAddress,
		moduleName:       moduleName,
		bankKeeper:       bankKeeper,
	}
}
