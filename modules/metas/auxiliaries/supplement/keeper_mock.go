/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/constants/test"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeperMock struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeperMock)(nil)

func (auxiliaryKeeper auxiliaryKeeperMock) Help(_ sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	var metaPropertyList []types.MetaProperty

	for _, property := range auxiliaryRequest.PropertyList {
		if property.GetID().String() == properties.Burn && property.GetFact().GetHashID().Equals(base.NewID("")) {
			return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList...), test.MockError)
		}
	}

	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(base.NewID(properties.Burn), base.NewMetaFact(base.NewHeightData(base.NewHeight(1)))))
	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewDecData(sdkTypes.SmallestDec()))))
	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(base.NewID(properties.TakerID), base.NewMetaFact(base.NewIDData(base.NewID("fromID")))))
	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(base.NewID(properties.ExchangeRate), base.NewMetaFact(base.NewDecData(sdkTypes.OneDec().Quo(sdkTypes.SmallestDec())))))
	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(base.NewID(properties.Expiry), base.NewMetaFact(base.NewHeightData(base.NewHeight(900)))))

	return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList...), nil)
}

func (auxiliaryKeeperMock) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeperMock{mapper: mapper}
}
func keeperPrototypeMock() helpers.AuxiliaryKeeper {
	return auxiliaryKeeperMock{}
}
