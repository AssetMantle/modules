// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/modules/metas/internal/mappable"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
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

		if property.GetHash().Compare(baseIDs.NewID("")) == 0 {
			if data, Error := baseTypes.ReadMetaProperty(property.GetType().String() + constants.DataTypeAndValueSeparator); Error == nil {
				meta = mappable.NewMeta(data.GetData())
			} else {
				return newAuxiliaryResponse(nil, Error)
			}
		} else {
			metaID := key.NewMetaID(property.GetType(), property.GetHash())
			metas := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.FromID(metaID))
			meta = metas.Get(key.FromID(metaID))
		}

		if meta != nil {
			metaPropertyList = append(metaPropertyList, baseTypes.NewMetaProperty(property.GetID(), meta.(mappables.Meta).GetData()))
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
