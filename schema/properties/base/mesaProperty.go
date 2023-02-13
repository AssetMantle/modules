// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ properties.MesaProperty = (*MesaProperty)(nil)

func (mesaProperty *MesaProperty) GetID() ids.PropertyID {
	return mesaProperty.ID
}
func (mesaProperty *MesaProperty) GetDataID() ids.DataID {
	return mesaProperty.DataID
}
func (mesaProperty *MesaProperty) GetKey() ids.StringID {
	return mesaProperty.ID.GetKey()
}
func (mesaProperty *MesaProperty) GetType() ids.StringID {
	return mesaProperty.ID.GetType()
}
func (mesaProperty *MesaProperty) GetBondWeight() int64 {
	if zeroData, err := base.PrototypeAnyData().FromString(mesaProperty.GetType().AsString()); err != nil {
		panic(err)
	} else {
		return zeroData.GetBondWeight()
	}
}
func (mesaProperty *MesaProperty) GetHash() ids.HashID {
	return mesaProperty.DataID.GetHashID()
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
func (mesaProperty *MesaProperty) ToAnyProperty() properties.AnyProperty {
	return &AnyProperty{
		Impl: &AnyProperty_MesaProperty{
			MesaProperty: mesaProperty,
		},
	}
}
func (mesaProperty *MesaProperty) ValidateBasic() error {
	// TODO implement
	return nil
}

func (mesaProperty *MesaProperty) Mutate(data data.Data) properties.Property {
	mesaProperty.DataID = data.GetID().(*baseIDs.DataID)
	return mesaProperty
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
		ID: propertyID.(*baseIDs.PropertyID),
	}
}
func NewMesaProperty(key ids.StringID, data data.Data) properties.MesaProperty {
	return &MesaProperty{
		ID:     baseIDs.NewPropertyID(key, data.GetType()).(*baseIDs.PropertyID),
		DataID: data.GetID().(*baseIDs.DataID),
	}
}
