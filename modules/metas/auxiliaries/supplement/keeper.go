// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/metas/module/key"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/utilities"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	propertyList := baseLists.NewPropertyList()

	for _, property := range auxiliaryRequest.PropertyList {
		if property.IsMeta() {
			propertyList = propertyList.Add(property)
		} else if property.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) == 0 {
			propertyList = propertyList.Add(baseProperties.NewMetaProperty(property.GetKey(), utilities.GetZeroValueDataFromID(property.GetType())))
		} else {
			metas := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(property.GetDataID()))
			if Mappable := metas.Get(key.NewKey(property.GetDataID())); Mappable != nil {
				propertyList = propertyList.Add(baseProperties.NewMetaProperty(property.GetKey(), Mappable.(data.Data)))
			}
		}
	}

	return newAuxiliaryResponse(propertyList, nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
