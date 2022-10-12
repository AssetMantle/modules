// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/metas/internal/mappable"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	scrubbedPropertyList := make([]properties.Property, len(auxiliaryRequest.MetaPropertyList))
	metas := auxiliaryKeeper.mapper.NewCollection(context)

	for i, metaProperty := range auxiliaryRequest.MetaPropertyList {
		if metaProperty.GetData().GenerateHashID().Compare(baseIDs.GenerateHashID()) != 0 {
			metas.Add(mappable.NewMeta(metaProperty.GetData()))
		}

		scrubbedPropertyList[i] = metaProperty.RemoveData()
	}

	return newAuxiliaryResponse(baseLists.NewPropertyList(scrubbedPropertyList...), nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
