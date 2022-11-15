// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/lists/base"
)

type auxiliaryKeeperMock struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeperMock)(nil)

func (auxiliaryKeeper auxiliaryKeeperMock) Help(_ sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	propertyList := base.NewPropertyList()

	for _, property := range auxiliaryRequest.PropertyList {
		//if property.GetID().Compare(constants.BurnHeightProperty.GetID()) == 0 && property.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) == 0 {
		//	return newAuxiliaryResponse(propertyList, errorConstants.MockError)
		//}
		if property.GetID().String() == "supplementError" {
			return newAuxiliaryResponse(nil, errorConstants.MockError)
		}
	}

	//propertyList = propertyList.Add(baseProperties.NewMetaProperty(constants.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))))
	//propertyList = propertyList.Add(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(sdkTypes.SmallestDec())))
	//propertyList = propertyList.Add(baseProperties.NewMetaProperty(constants.TakerIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewStringID("fromID"))))
	//propertyList = propertyList.Add(baseProperties.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(sdkTypes.OneDec().Quo(sdkTypes.SmallestDec()))))
	//propertyList = propertyList.Add(baseProperties.NewMetaProperty(constants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(900))))

	return newAuxiliaryResponse(propertyList, nil)
}

func (auxiliaryKeeperMock) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeperMock{mapper: mapper}
}
func keeperPrototypeMock() helpers.AuxiliaryKeeper {
	return auxiliaryKeeperMock{}
}
