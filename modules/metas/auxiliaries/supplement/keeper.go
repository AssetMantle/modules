/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/metas/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	var metaPropertyList []types.MetaProperty
	for _, property := range auxiliaryRequest.Properties.GetList() {
		metaID := base.NewID(property.GetFact().GetHash())
		metas := mapper.NewMetas(auxiliaryKeeper.mapper, context).Fetch(metaID)
		meta := metas.Get(metaID)
		if meta != nil {
			metaPropertyList = append(metaPropertyList, base.NewMetaProperty(property.GetID(), base.NewMetaFact(meta.GetData())))
		}
	}
	return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList), nil)
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
