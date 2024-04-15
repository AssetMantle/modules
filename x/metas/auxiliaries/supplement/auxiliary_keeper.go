// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"context"
	"github.com/AssetMantle/schema/go/errors/constants"
	"reflect"

	"github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/metas/key"
	"github.com/AssetMantle/modules/x/metas/mappable"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, AuxiliaryRequest helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest, ok := AuxiliaryRequest.(auxiliaryRequest)
	if !ok {
		return nil, constants.InvalidRequest.Wrapf("invalid request type: %s", reflect.TypeOf(AuxiliaryRequest).String())
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

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
