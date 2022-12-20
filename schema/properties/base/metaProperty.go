// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

//type metaProperty struct {
//	ID   ids.PropertyID `json:"id"`
//	Data data.Data      `json:"data"`
//}

var _ properties.MetaProperty = (*Property_MetaProperty)(nil)

func (metaProperty *Property_MetaProperty) GetData() data.Data {
	return metaProperty.MetaProperty.Data
}
func (metaProperty *Property_MetaProperty) ScrubData() properties.Property {
	return NewMesaProperty(metaProperty.GetKey(), metaProperty.GetData())
}
func (metaProperty *Property_MetaProperty) GetID() ids.ID {
	return metaProperty.MetaProperty.ID
}
func (metaProperty *Property_MetaProperty) GetDataID() ids.ID {
	return metaProperty.MetaProperty.Data.GetID()
}
func (metaProperty *Property_MetaProperty) GetKey() ids.ID {
	return metaProperty.MetaProperty.ID.Impl.(properties.Property).GetKey()
}
func (metaProperty *Property_MetaProperty) GetType() ids.ID {
	return metaProperty.MetaProperty.Data.GetType()
}
func (metaProperty *Property_MetaProperty) IsMeta() bool {
	return true
}
func (metaProperty *Property_MetaProperty) Compare(listable traits.Listable) int {
	// NOTE: compare property can be meta or mesa, so listable must only be cast to Property Interface
	if compareProperty, err := propertyFromInterface(listable); err != nil {
		panic(err)
	} else {
		return metaProperty.GetID().Compare(compareProperty.GetID())
	}
}

func NewEmptyMetaPropertyFromID(propertyID ids.ID) properties.Property {
	return &Property{
		Impl: &Property_MetaProperty{
			MetaProperty: &MetaProperty{
				ID: propertyID.(*baseIDs.ID),
			},
		},
	}
}
func NewMetaProperty(key ids.StringID, data data.Data) properties.Property {
	if data == nil || key == nil {
		panic(errorConstants.MetaDataError)
	}
	return &Property{
		Impl: &Property_MetaProperty{
			MetaProperty: &MetaProperty{
				ID:   baseIDs.NewPropertyID(key, data.GetType()).(*baseIDs.ID),
				Data: data.(*baseData.Data),
			},
		},
	}
}
