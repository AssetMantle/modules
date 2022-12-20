// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// type mesaProperty struct {
//	ID     ids.PropertyID
//	DataID ids.DataID
// }

var _ properties.MesaProperty = (*MesaProperty)(nil)

func (mesaProperty *MesaProperty) ScrubData() properties.Property {
	panic("this method should never be called")
}

func (mesaProperty *MesaProperty) GetID() ids.PropertyID {
	return mesaProperty.Id
}
func (mesaProperty *MesaProperty) GetDataID() ids.DataID {
	return mesaProperty.DataId
}
func (mesaProperty *MesaProperty) GetKey() ids.StringID {
	return mesaProperty.Id.GetKey()
}
func (mesaProperty *MesaProperty) GetType() ids.StringID {
	return mesaProperty.Id.GetType()
}
func (mesaProperty *MesaProperty) GetHash() ids.HashID {
	return mesaProperty.DataId.GetHashID()
}
func (mesaProperty *MesaProperty) IsMeta() bool {
	return false
}
func (mesaProperty *MesaProperty) IsMesa() bool {
	return true
}
func (mesaProperty *MesaProperty) Compare(listable traits.Listable) int {
	// NOTE: compare property can be meta or mesa, so listable must only be cast to Property Interface and not MesaProperty
	if compareProperty, err := propertyFromInterface(listable); err != nil {
		panic(err)
	} else {
		return mesaProperty.GetID().Compare(compareProperty.GetID())
	}
}
func (mesaProperty *MesaProperty) GetData() data.AnyData {
	panic("This is meant to be unreachable")
}
func (mesaProperty *MesaProperty) ToAnyProperty() properties.AnyProperty {
	return &AnyProperty{
		Impl: &AnyProperty_MesaProperty{
			MesaProperty: mesaProperty,
		},
	}
}
func propertyFromInterface(listable traits.Listable) (properties.Property, error) {
	switch value := listable.(type) {
	case properties.Property:
		return value, nil
	default:
		return nil, constants.MetaDataError
	}
}
func NewEmptyMesaPropertyFromID(propertyID ids.PropertyID) properties.MesaProperty {
	return &MesaProperty{
		Id: propertyID.(*baseIDs.PropertyID),
	}
}
func NewMesaProperty(key ids.StringID, data data.Data) properties.Property {
	return &MesaProperty{
		Id:     baseIDs.NewPropertyID(key, data.GetType()).(*baseIDs.PropertyID),
		DataId: data.GetID().(*baseIDs.DataID),
	}
}
