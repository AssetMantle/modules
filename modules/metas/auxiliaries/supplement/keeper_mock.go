/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeperMock struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeperMock)(nil)

func (auxiliaryKeeper auxiliaryKeeperMock) Help(_ sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	var metaPropertyList []types.MetaProperty
	for _, property := range auxiliaryRequest.PropertyList {
		if property.GetID().String() == "burn" && property.GetFact().GetHash() == "" {
			return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList), errors.MockError)
		}
	}
	metaPropertyList = append(metaPropertyList, base.NewMetaProperty(base.NewID(properties.Burn), base.NewMetaFact(base.NewHeightData(base.NewHeight(1)))))
	return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList), nil)
}

func initializeAuxiliaryKeeperMock(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeperMock{mapper: mapper}
}
