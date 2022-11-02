// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
)

type auxiliaryKeeperMock struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeperMock)(nil)

func (auxiliaryKeeper auxiliaryKeeperMock) Help(_ sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	scrubbedPropertyList := make([]properties.Property, len(auxiliaryRequest.PropertyList.GetList()))

	for i, property := range auxiliaryRequest.PropertyList.GetList() {
		if property.IsMeta() {
			scrubbedPropertyList[i] = property.(properties.MetaProperty).ScrubData()
		} else {
			scrubbedPropertyList[i] = property
		}

		if property.GetID().String() == "scrubError" {
			return newAuxiliaryResponse(nil, constants.MockError)
		}
	}

	return newAuxiliaryResponse(baseLists.NewPropertyList(scrubbedPropertyList...), nil)
}

func (auxiliaryKeeperMock) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeperMock{mapper: mapper}
}
func keeperPrototypeMock() helpers.AuxiliaryKeeper {
	return auxiliaryKeeperMock{}
}
