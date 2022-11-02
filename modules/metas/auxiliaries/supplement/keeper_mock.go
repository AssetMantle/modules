// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	base2 "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type auxiliaryKeeperMock struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeperMock)(nil)

func (auxiliaryKeeper auxiliaryKeeperMock) Help(_ sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	propertyList := base.NewPropertyList()

	for _, property := range auxiliaryRequest.PropertyList {
		if property.GetID().Compare(constants.BurnHeightProperty.GetID()) == 0 && property.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) == 0 {
			return newAuxiliaryResponse(propertyList, errorConstants.MockError)
		}
	}

	propertyList = propertyList.Add(base2.NewMetaProperty(constants.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))))
	propertyList = propertyList.Add(base2.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(sdkTypes.SmallestDec())))
	propertyList = propertyList.Add(base2.NewMetaProperty(constants.TakerIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewStringID("fromID"))))
	propertyList = propertyList.Add(base2.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(sdkTypes.OneDec().Quo(sdkTypes.SmallestDec()))))
	propertyList = propertyList.Add(base2.NewMetaProperty(constants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(900))))

	return newAuxiliaryResponse(propertyList, nil)
}

func (auxiliaryKeeperMock) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeperMock{mapper: mapper}
}
func keeperPrototypeMock() helpers.AuxiliaryKeeper {
	return auxiliaryKeeperMock{}
}
