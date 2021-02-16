/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
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

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	var metaPropertyList []types.MetaProperty

	for _, property := range auxiliaryRequest.PropertyList {
		var meta helpers.Mappable

		if property.GetFact().GetHashID().Equals(base.NewID("")) {
			if metaFact, Error := base.ReadMetaFact(property.GetFact().GetTypeID().String() + constants.DataTypeAndValueSeparator); Error == nil {
				meta = mappable.NewMeta(metaFact.GetData())
			} else {
				return newAuxiliaryResponse(nil, Error)
			}
		} else {
			metaID := key.NewMetaID(property.GetFact().GetTypeID(), property.GetFact().GetHashID())
			metas := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.FromID(metaID))
			meta = metas.Get(key.FromID(metaID))
		}

		if meta != nil {
			metaPropertyList = append(metaPropertyList, base.NewMetaProperty(property.GetID(), base.NewMetaFact(meta.(mappables.Meta).GetData())))
		}
	}

	return newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList...), nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
