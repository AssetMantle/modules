// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"context"
	"github.com/AssetMantle/modules/helpers/constants"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/metas/record"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, AuxiliaryRequest helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest, ok := AuxiliaryRequest.(auxiliaryRequest)
	if !ok {
		return nil, constants.InvalidRequest.Wrapf("invalid request type %T", AuxiliaryRequest)
	}

	if err := auxiliaryRequest.Validate(); err != nil {
		return nil, err
	}

	scrubbedPropertyList := make([]properties.Property, len(auxiliaryRequest.PropertyList.Get()))
	metas := auxiliaryKeeper.mapper.NewCollection(context)

	for i, property := range auxiliaryRequest.PropertyList.Get() {
		if property.IsMeta() {
			metaProperty := property.Get().(properties.MetaProperty)
			if metaProperty.GetData().GenerateHashID().Compare(baseIDs.GenerateHashID()) != 0 {

				if err := metaProperty.GetData().ValidateBasic(); err != nil {
					return nil, err
				}

				metas.Add(record.NewRecord(metaProperty.GetData()))
			}
			scrubbedPropertyList[i] = metaProperty.ScrubData()
		} else {
			scrubbedPropertyList[i] = property
		}
	}

	return newAuxiliaryResponse(baseLists.NewPropertyList(scrubbedPropertyList...)), nil
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper

	helpers.PanicOnUninitializedKeeperFields(auxiliaryKeeper)
	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
