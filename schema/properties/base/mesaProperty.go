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
	ID     ids.PropertyID
	DataID ids.DataID
}

var _ properties.MesaProperty = (*mesaProperty)(nil)

func (mesaProperty mesaProperty) GetID() ids.PropertyID {
	return mesaProperty.ID
}
func (mesaProperty mesaProperty) GetDataID() ids.DataID {
	return mesaProperty.DataID
}
func (mesaProperty mesaProperty) GetKey() ids.StringID {
	return mesaProperty.ID.GetKey()
}
func (mesaProperty mesaProperty) GetType() ids.StringID {
	return mesaProperty.ID.GetType()
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
	if compareProperty, err := mesaPropertyFromInterface(listable); err != nil {
		panic(err)
	} else {
		return mesaProperty.GetID().Compare(compareProperty.GetID())
	}
}
func mesaPropertyFromInterface(listable traits.Listable) (mesaProperty, error) {
	switch value := listable.(type) {
	case mesaProperty:
		return value, nil
	default:
		return mesaProperty{}, constants.MetaDataError
	}
}

func NewEmptyMesaPropertyFromID(propertyID ids.PropertyID) properties.Property {
	return mesaProperty{
		ID: propertyID,
	}
}
func NewMesaProperty(key ids.StringID, data data.Data) properties.Property {
	return mesaProperty{
		ID:     baseIDs.NewPropertyID(key, data.GetType()),
		DataID: data.GetID(),
	}
}
