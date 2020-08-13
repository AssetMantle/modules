/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package reverse

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type auxiliaryRequest struct {
	MakerID      types.ID
	MakerSplit   sdkTypes.Dec
	MakerSplitID types.ID
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

func NewAuxiliaryRequest(makerID types.ID, makerSplit sdkTypes.Dec, makerSplitID types.ID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		MakerID:      makerID,
		MakerSplit:   makerSplit,
		MakerSplitID: makerSplitID,
	}
}
