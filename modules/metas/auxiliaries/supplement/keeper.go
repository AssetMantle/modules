// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"context"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/modules/metas/internal/mappable"
	"github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	propertyList := baseLists.NewPropertyList()

	for _, property := range auxiliaryRequest.PropertyList {
		if property.IsMeta() {
			propertyList = propertyList.Add(property)
		} else if property.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) == 0 {
			if zeroData, err := base.PrototypeAnyData().FromString(property.GetDataID().AsString()); err == nil {
				propertyList = propertyList.Add(baseProperties.NewMetaProperty(property.GetKey(), zeroData))
			}
		} else {
			metas := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(property.GetDataID()))
			if Mappable := metas.Get(key.NewKey(property.GetDataID())); Mappable != nil {
				propertyList = propertyList.Add(baseProperties.NewMetaProperty(property.GetKey(), mappable.GetData(Mappable)))
			}
		}
	}

	return newAuxiliaryResponse(propertyList, nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterList, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
