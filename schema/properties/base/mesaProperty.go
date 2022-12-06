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

type mesaProperty struct {
	ids.PropertyID
	ids.DataID
}

var _ properties.MesaProperty = (*mesaProperty)(nil)

func (mesaProperty mesaProperty) GetID() ids.PropertyID {
	return mesaProperty.PropertyID
}
func (mesaProperty mesaProperty) GetDataID() ids.DataID {
	return mesaProperty.DataID
}
func (mesaProperty mesaProperty) GetKey() ids.StringID {
	return mesaProperty.PropertyID.GetKey()
}
func (mesaProperty mesaProperty) GetType() ids.StringID {
	return mesaProperty.PropertyID.GetType()
}
func (mesaProperty mesaProperty) GetHash() ids.ID {
	return mesaProperty.DataID.GetHashID()
}
func (mesaProperty mesaProperty) IsMeta() bool {
	return false
}
func (mesaProperty mesaProperty) IsMesa() bool {
	return true
}
func (mesaProperty mesaProperty) Compare(listable traits.Listable) int {
	// NOTE: compare property can be meta or mesa, so listable must only be cast to Property Interface and not MesaProperty
	if compareProperty, err := propertyFromInterface(listable); err != nil {
		panic(err)
	} else {
		return mesaProperty.GetID().Compare(compareProperty.GetID())
	}
}
func propertyFromInterface(listable traits.Listable) (properties.Property, error) {
	switch value := listable.(type) {
	case properties.Property:
		return value, nil
	default:
		return mesaProperty{}, constants.MetaDataError
	}
}

func NewEmptyMesaPropertyFromID(propertyID ids.PropertyID) properties.Property {
	return mesaProperty{
		PropertyID: propertyID,
	}
}
func NewMesaProperty(key ids.StringID, data data.Data) properties.MesaProperty {
	return mesaProperty{
		PropertyID: baseIDs.GeneratePropertyID(key, data.GetType()),
		DataID:     data.GetID(),
	}
}
