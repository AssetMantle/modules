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
//type mesaProperty struct {
//	ID     ids.PropertyID
//	DataID ids.DataID
//}

var _ properties.MesaProperty = (*Property_MesaProperty)(nil)

func (mesaProperty *Property_MesaProperty) ScrubData() properties.Property {
	panic("this method should never be called")
}

func (mesaProperty *Property_MesaProperty) GetID() ids.ID {
	return mesaProperty.MesaProperty.ID
}
func (mesaProperty *Property_MesaProperty) GetDataID() ids.ID {
	return mesaProperty.MesaProperty.DataID
}
func (mesaProperty *Property_MesaProperty) GetKey() ids.ID {
	return mesaProperty.MesaProperty.ID.Impl.(ids.PropertyID).GetKey()
}
func (mesaProperty *Property_MesaProperty) GetType() ids.ID {
	return mesaProperty.MesaProperty.ID.Impl.(ids.PropertyID).GetType()
}
func (mesaProperty *Property_MesaProperty) GetHash() ids.ID {
	return mesaProperty.MesaProperty.DataID.Impl.(ids.DataID).GetHashID()
}
func (mesaProperty *Property_MesaProperty) IsMeta() bool {
	return false
}
func (mesaProperty *Property_MesaProperty) IsMesa() bool {
	return true
}
func (mesaProperty *Property_MesaProperty) Compare(listable traits.Listable) int {
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
		return nil, constants.MetaDataError
	}
}

func NewEmptyMesaPropertyFromID(propertyID ids.ID) properties.Property {
	return &Property{
		Impl: &Property_MesaProperty{
			MesaProperty: &MesaProperty{
				ID: propertyID.(*baseIDs.ID),
			},
		},
	}
}
func NewMesaProperty(key ids.ID, data data.Data) properties.Property {
	return &Property{
		Impl: &Property_MesaProperty{
			MesaProperty: &MesaProperty{
				ID:     baseIDs.NewPropertyID(key, data.GetType()).(*baseIDs.ID),
				DataID: data.GetID().(*baseIDs.ID),
			},
		},
	}
}
