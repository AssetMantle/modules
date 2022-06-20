// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

type property struct {
	ID     ids.PropertyID `json:"id"`
	DataID ids.DataID     `json:"dataID"`
}

var _ properties.Property = (*property)(nil)

func (property property) GetID() ids.PropertyID {
	return property.ID
}
func (property property) GetDataID() ids.DataID {
	return property.DataID
}
func (property property) GetKey() ids.ID {
	return property.ID.GetKey()
}
func (property property) GetType() ids.ID {
	return property.ID.GetType()
}
func (property property) GetHash() ids.ID {
	return property.DataID.GetHash()
}
func (property property) Compare(listable traits.Listable) int {
	if compareProperty, err := propertyFromInterface(listable); err != nil {
		panic(err)
	} else {
		return property.GetID().Compare(compareProperty.GetID())
	}
}
func propertyFromInterface(listable traits.Listable) (property, error) {
	switch value := listable.(type) {
	case property:
		return value, nil
	default:
		return property{}, constants.MetaDataError
	}
}

func NewEmptyPropertyFromID(propertyID ids.PropertyID) properties.Property {
	return property{
		ID: propertyID,
	}
}
func NewPropertyWithDataID(propertyID ids.PropertyID, dataID ids.DataID) properties.Property {
	return property{
		ID:     propertyID,
		DataID: dataID,
	}
}
func NewProperty(key ids.ID, data data.Data) properties.Property {
	return property{
		ID:     baseIDs.NewPropertyID(key, data.GetType()),
		DataID: data.GetID(),
	}
}
