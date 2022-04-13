// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

type property struct {
	ID     ids.PropertyID `json:"id"`
	DataID ids.DataID     `json:"dataID"`
}

var _ types.Property = (*property)(nil)

func (property property) GetID() ids.PropertyID {
	return property.ID
}
func (property property) GetDataID() ids.DataID {
	return property.DataID
}
func (property property) GetKey() types.ID {
	return property.ID.GetKey()
}
func (property property) GetType() types.ID {
	return property.ID.GetType()
}
func (property property) GetHash() types.ID {
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
		return property{}, errors.MetaDataError
	}
}

func NewEmptyPropertyFromID(propertyID ids.PropertyID) types.Property {
	return property{
		ID: propertyID,
	}
}
func NewProperty(key types.ID, data types.Data) types.Property {
	return property{
		ID:     baseIDs.NewPropertyID(key, data.GetType()),
		DataID: data.GetID(),
	}
}
