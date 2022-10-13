// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package member

import (
	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/mappables"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
)

type auxiliaryKeeperMock struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeperMock)(nil)

func (auxiliaryKeeper auxiliaryKeeperMock) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.ClassificationID))
	Mappable := classifications.Get(key.NewKey(auxiliaryRequest.ClassificationID))
	if Mappable == nil {
		return newAuxiliaryResponse(errorConstants.EntityNotFound)
	}
	classification := Mappable.(mappables.Classification)
	if auxiliaryRequest.Mutables != nil {
		if len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()) > len(classification.GetMutables().GetMutablePropertyList().GetList()) {
			return newAuxiliaryResponse(errorConstants.IncorrectFormat)
		}

		for _, mutableProperty := range auxiliaryRequest.Mutables.GetMutablePropertyList().GetList() {
			if property := classification.GetMutables().GetMutablePropertyList().GetProperty(mutableProperty.GetID()); property == nil {
				return newAuxiliaryResponse(errorConstants.IncorrectFormat)
			}
		}
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeperMock) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeperMock{mapper: mapper}
}
func keeperPrototypeMock() helpers.AuxiliaryKeeper {
	return auxiliaryKeeperMock{}
}
