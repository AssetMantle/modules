// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/constants/test"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

type auxiliaryKeeperMock struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeperMock)(nil)

func (auxiliaryKeeper auxiliaryKeeperMock) Help(_ sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	var metaPropertyList []types.MetaProperty

	for _, property := range auxiliaryRequest.PropertyList {
		if property.GetID().Compare(ids.BurnProperty) == 0 && property.GetHashID().Compare(base.NewID("")) == 0 {
			return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList...), test.MockError)
		}
	}

	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(ids.BurnProperty, base.NewHeightData(base.NewHeight(1))))
	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(ids.MakerOwnableSplitProperty, base.NewDecData(sdkTypes.SmallestDec())))
	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(ids.TakerIDProperty, base.NewIDData(base.NewID("fromID"))))
	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(ids.ExchangeRateProperty, base.NewDecData(sdkTypes.OneDec().Quo(sdkTypes.SmallestDec()))))
	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(ids.ExpiryProperty, base.NewHeightData(base.NewHeight(900))))

	return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList...), nil)
}

func (auxiliaryKeeperMock) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeperMock{mapper: mapper}
}
func keeperPrototypeMock() helpers.AuxiliaryKeeper {
	return auxiliaryKeeperMock{}
}
