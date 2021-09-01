/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package scrub

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	scrubbedPropertyList := make([]base.Property, len(auxiliaryRequest.MetaPropertyList))
	metas := auxiliaryKeeper.mapper.NewCollection(context)

	for i, metaProperty := range auxiliaryRequest.MetaPropertyList {
		if metaProperty.GetMetaFact().GetHashID().Compare(base.NewID("")) != 0 {
			metas.Add(mappable.NewMeta(metaProperty.GetMetaFact().GetData()))
		}
		a := base.Property{
			Id:   *base.NewID(metaProperty.GetID().String()),
			Fact: *base.NewFact(metaProperty.GetMetaFact().GetData()),
		}
		scrubbedPropertyList[i] = a
	}

	return newAuxiliaryResponse(base.NewProperties(scrubbedPropertyList...), nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
