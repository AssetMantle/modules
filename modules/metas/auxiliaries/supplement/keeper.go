// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/modules/metas/internal/mappable"
	"github.com/AssetMantle/modules/schema/data/utilities"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	var metaPropertyList []properties.MetaProperty

	for _, property := range auxiliaryRequest.PropertyList {
		var meta helpers.Mappable

		if property.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) == 0 {
			meta = mappable.NewMeta(utilities.GetZeroValueDataFromID(property.GetType()))
		} else {
			metaID := baseIDs.NewMetaID(property.GetType(), property.GetDataID().GetHashID())
			metas := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(metaID))
			meta = metas.Get(key.NewKey(metaID))
		}

		if meta != nil {
			metaPropertyList = append(metaPropertyList, baseProperties.NewMetaProperty(property.GetKey(), meta.(mappables.Meta).GetData()))
		}
	}

	return newAuxiliaryResponse(base.NewMetaPropertyList(metaPropertyList...), nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
