// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/test"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
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

	var metaPropertyList []properties.MetaProperty

	for _, property := range auxiliaryRequest.PropertyList {
		if property.GetID().Compare(constants.BurnProperty) == 0 && property.GetHash().Compare(baseIDs.NewID("")) == 0 {
			return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList...), test.MockError)
		}
	}

	metaPropertyList = append(metaPropertyList, base2.NewMetaProperty(constants.BurnProperty, baseData.NewHeightData(baseTypes.NewHeight(1))))
	metaPropertyList = append(metaPropertyList, base2.NewMetaProperty(constants.MakerOwnableSplitProperty, baseData.NewDecData(sdkTypes.SmallestDec())))
	metaPropertyList = append(metaPropertyList, base2.NewMetaProperty(constants.TakerIDProperty, baseData.NewIDData(baseIDs.NewID("fromID"))))
	metaPropertyList = append(metaPropertyList, base2.NewMetaProperty(constants.ExchangeRateProperty, baseData.NewDecData(sdkTypes.OneDec().Quo(sdkTypes.SmallestDec()))))
	metaPropertyList = append(metaPropertyList, base2.NewMetaProperty(constants.ExpiryProperty, baseData.NewHeightData(baseTypes.NewHeight(900))))

	return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList...), nil)
}

func (auxiliaryKeeperMock) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeperMock{mapper: mapper}
}
func keeperPrototypeMock() helpers.AuxiliaryKeeper {
	return auxiliaryKeeperMock{}
}
