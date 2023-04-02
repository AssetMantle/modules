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

var _ properties.MetaProperty = (*MetaProperty)(nil)

func (metaProperty *MetaProperty) IsMetaProperty() {
}
func (metaProperty *MetaProperty) ValidateBasic() error {
	if err := metaProperty.ID.ValidateBasic(); err != nil {
		return err
	}
	if err := metaProperty.Data.ValidateBasic(); err != nil {
		return err
	}
	if metaProperty.Data.GetTypeID().Compare(metaProperty.ID.TypeID) != 0 {
		return errorConstants.IncorrectFormat.Wrapf("data type id does not match property type id")
	}
	return nil
}
func (metaProperty *MetaProperty) GetData() data.AnyData {
	return metaProperty.Data
}
func (metaProperty *MetaProperty) ScrubData() properties.MesaProperty {
	return NewMesaProperty(metaProperty.GetKey(), metaProperty.GetData())
}
func (metaProperty *MetaProperty) GetID() ids.PropertyID {
	return metaProperty.ID
}
func (metaProperty *MetaProperty) GetDataID() ids.DataID {
	return metaProperty.Data.GetID()
}
func (metaProperty *MetaProperty) GetKey() ids.StringID {
	return metaProperty.ID.GetKey()
}
func (metaProperty *MetaProperty) GetDataTypeID() ids.StringID {
	return metaProperty.ID.GetDataTypeID()
}
func (metaProperty *MetaProperty) GetBondWeight() int64 {
	return metaProperty.Data.GetBondWeight()
}
func (metaProperty *MetaProperty) IsMeta() bool {
	return true
}
func (metaProperty *MetaProperty) Mutate(data data.Data) properties.Property {
	metaProperty.Data = data.ToAnyData().(*base.AnyData)
	return metaProperty
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
		ID: propertyID.(*baseIDs.PropertyID),
	}
}
func NewMetaProperty(key ids.StringID, data data.Data) properties.MetaProperty {
	if data == nil || key == nil {
		panic(errorConstants.IncorrectFormat.Wrapf("meta property data or key cannot be nil"))
	}
	return &MetaProperty{
		ID:   baseIDs.NewPropertyID(key, data.GetTypeID()).(*baseIDs.PropertyID),
		Data: data.ToAnyData().(*base.AnyData),
	}
}
