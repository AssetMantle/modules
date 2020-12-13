/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mappable"
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
		var meta helpers.Mappable
		if property.GetFact().GetHash() == "" {
			if metaFact, Error := base.ReadMetaFact(property.GetFact().GetType() + constants.DataTypeAndValueSeparator); Error == nil {
				meta = mappable.NewMeta(metaFact.GetData())
			} else {
				return newAuxiliaryResponse(nil, Error)
			}
		} else {
			metaID := base.NewID(property.GetFact().GetHash())
			metas := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.New(metaID))
			meta = metas.Get(key.New(metaID))
		}
		if meta != nil {
			metaPropertyList = append(metaPropertyList, base.NewMetaProperty(property.GetID(), base.NewMetaFact(meta.(mappables.Meta).GetData())))
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
