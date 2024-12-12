// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/metas/key"
	"github.com/AssetMantle/modules/x/metas/mappable"
	"github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, AuxiliaryRequest helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest, ok := AuxiliaryRequest.(auxiliaryRequest)
	if !ok {
		return nil, errorConstants.InvalidRequest.Wrapf("invalid request type %T", AuxiliaryRequest)
	}

	if err := auxiliaryRequest.Validate(); err != nil {
		return nil, err
	}

	propertyList := baseLists.NewPropertyList()

	for _, property := range auxiliaryRequest.PropertyList {
		if property == nil {
			continue
		} else if property.IsMeta() {
			propertyList = propertyList.Add(property)
		} else if property.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) == 0 {
			if zeroData, err := base.PrototypeAnyData().FromString(property.GetDataTypeID().AsString()); err == nil {
				propertyList = propertyList.Add(baseProperties.NewMetaProperty(property.GetKey(), zeroData))
			}
		} else {
			metas := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(property.GetDataID()))
			if Mappable := metas.GetMappable(key.NewKey(property.GetDataID())); Mappable != nil {
				propertyList = propertyList.Add(baseProperties.NewMetaProperty(property.GetKey(), mappable.GetData(Mappable)))
			}
		}
	}

	return NewAuxiliaryResponse(propertyList), nil
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper

	helpers.PanicOnUninitializedKeeperFields(auxiliaryKeeper)
	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
