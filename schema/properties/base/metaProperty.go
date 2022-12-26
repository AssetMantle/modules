// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

// type metaProperty struct {
//	ID   ids.PropertyID `json:"id"`
//	Data data.Data      `json:"data"`
// }

var _ properties.MetaProperty = (*MetaProperty)(nil)

func (metaProperty *MetaProperty) GetData() data.AnyData {
	return metaProperty.AnyData
}
func (metaProperty *MetaProperty) ScrubData() properties.Property {
	return NewMesaProperty(metaProperty.GetKey(), metaProperty.GetData())
}
func (metaProperty *MetaProperty) GetID() ids.PropertyID {
	return metaProperty.Id
}
func (metaProperty *MetaProperty) GetDataID() ids.DataID {
	return metaProperty.AnyData.GetID()
}
func (metaProperty *MetaProperty) GetKey() ids.StringID {
	return metaProperty.Id.GetKey()
}
func (metaProperty *MetaProperty) GetType() ids.StringID {
	return metaProperty.AnyData.GetType()
}
func (metaProperty *MetaProperty) IsMeta() bool {
	return true
}
func (metaProperty *MetaProperty) Compare(listable traits.Listable) int {
	// NOTE: compare property can be meta or mesa, so listable must only be cast to Property Interface
	if compareProperty, err := propertyFromInterface(listable); err != nil {
		panic(err)
	} else {
		return metaProperty.GetID().Compare(compareProperty.GetID())
	}
}
func (metaProperty *MetaProperty) ToAnyProperty() properties.AnyProperty {
	return &AnyProperty{
		Impl: &AnyProperty_MetaProperty{
			MetaProperty: metaProperty,
		},
	}
}
func NewEmptyMetaPropertyFromID(propertyID ids.PropertyID) properties.MetaProperty {
	return &MetaProperty{
		Id: propertyID.(*baseIDs.PropertyID),
	}
}
func NewMetaProperty(key ids.StringID, data data.AnyData) properties.Property {
	if data == nil || key == nil {
		panic(errorConstants.MetaDataError)
	}
	return &MetaProperty{
		Id:      baseIDs.NewPropertyID(key, data.GetType()).(*baseIDs.PropertyID),
		AnyData: data.ToAnyData().(*base.AnyData),
	}
}
