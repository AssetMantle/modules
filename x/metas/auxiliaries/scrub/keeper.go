// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"context"

	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	"github.com/AssetMantle/schema/x/properties"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/metas/internal/mappable"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	scrubbedPropertyList := make([]properties.Property, len(auxiliaryRequest.PropertyList.GetList()))
	metas := auxiliaryKeeper.mapper.NewCollection(context)

	for i, property := range auxiliaryRequest.PropertyList.GetList() {
		if property.IsMeta() {
			metaProperty := property.Get().(properties.MetaProperty)
			if metaProperty.GetData().GenerateHashID().Compare(baseIDs.GenerateHashID()) != 0 {
				metas.Add(mappable.NewMappable(metaProperty.GetData()))
			}
			scrubbedPropertyList[i] = metaProperty.ScrubData()
		} else {
			scrubbedPropertyList[i] = property
		}
	}

	return newAuxiliaryResponse(baseLists.NewPropertyList(scrubbedPropertyList...)), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
