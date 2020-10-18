/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
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
	for _, property := range auxiliaryRequest.PropertyList {
		var meta mappables.Meta
		if property.GetFact().GetHash() == "" {
			var data types.Data
			switch property.GetFact().GetType() {
			case base.DecType:
				data, _ = base.ReadDecData("")
			case base.HeightType:
				data, _ = base.ReadHeightData("")
			case base.IDType:
				data, _ = base.ReadIDData("")
			case base.StringType:
				data, _ = base.ReadStringData("")
			}
			meta = mappable.NewMeta(data)
		} else {
			metaID := base.NewID(property.GetFact().GetHash())
			metas := mapper.NewMetas(auxiliaryKeeper.mapper, context).Fetch(metaID)
			meta = metas.Get(metaID)
		}
		if meta != nil {
			metaPropertyList = append(metaPropertyList, base.NewMetaProperty(property.GetID(), base.NewMetaFact(meta.GetData())))
		}
	}
	return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList), nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
