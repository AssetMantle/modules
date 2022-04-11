// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/constants/test"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type auxiliaryKeeperMock struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeperMock)(nil)

func (auxiliaryKeeper auxiliaryKeeperMock) Help(_ sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	var metaPropertyList []types.MetaProperty

	for _, property := range auxiliaryRequest.PropertyList {
		if property.GetID().Compare(ids.BurnProperty) == 0 && property.GetHash().Compare(baseIDs.NewID("")) == 0 {
			return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList...), test.MockError)
		}
	}

	metaPropertyList = append(metaPropertyList, baseTypes.NewMetaProperty(ids.BurnProperty, baseData.NewHeightData(baseTypes.NewHeight(1))))
	metaPropertyList = append(metaPropertyList, baseTypes.NewMetaProperty(ids.MakerOwnableSplitProperty, baseData.NewDecData(sdkTypes.SmallestDec())))
	metaPropertyList = append(metaPropertyList, baseTypes.NewMetaProperty(ids.TakerIDProperty, baseData.NewIDData(baseIDs.NewID("fromID"))))
	metaPropertyList = append(metaPropertyList, baseTypes.NewMetaProperty(ids.ExchangeRateProperty, baseData.NewDecData(sdkTypes.OneDec().Quo(sdkTypes.SmallestDec()))))
	metaPropertyList = append(metaPropertyList, baseTypes.NewMetaProperty(ids.ExpiryProperty, baseData.NewHeightData(baseTypes.NewHeight(900))))

	return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList...), nil)
}

func (auxiliaryKeeperMock) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeperMock{mapper: mapper}
}
func keeperPrototypeMock() helpers.AuxiliaryKeeper {
	return auxiliaryKeeperMock{}
}
