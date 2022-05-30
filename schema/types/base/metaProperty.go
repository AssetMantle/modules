// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data/utilities"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

type metaProperty struct {
	ID   ids.PropertyID `json:"id"`
	Data types.Data     `json:"data"`
}

var _ types.MetaProperty = (*metaProperty)(nil)

func (metaProperty metaProperty) GetData() types.Data {
	return metaProperty.Data
}
func (metaProperty metaProperty) RemoveData() types.Property {
	return NewProperty(metaProperty.GetKey(), metaProperty.GetData())
}
func (metaProperty metaProperty) GetID() ids.PropertyID {
	return metaProperty.ID
}
func (metaProperty metaProperty) GetDataID() ids.DataID {
	return metaProperty.Data.GetID()
}
func (metaProperty metaProperty) GetKey() types.ID {
	return metaProperty.ID.GetKey()
}
func (metaProperty metaProperty) GetType() types.ID {
	return metaProperty.Data.GetType()
}
func (metaProperty metaProperty) GetHash() types.ID {
	return metaProperty.Data.GenerateHash()
}
func (metaProperty metaProperty) Compare(listable traits.Listable) int {
	if compareMetaProperty, err := metaPropertyFromInterface(listable); err != nil {
		panic(err)
	} else {
		return metaProperty.GetID().Compare(compareMetaProperty.GetID())
	}
}
func metaPropertyFromInterface(listable traits.Listable) (metaProperty, error) {
	switch value := listable.(type) {
	case metaProperty:
		return value, nil
	default:
		return metaProperty{}, errors.MetaDataError
	}
}

func NewEmptyMetaPropertyFromID(propertyID ids.PropertyID) types.MetaProperty {
	return metaProperty{
		ID: propertyID,
	}
}
func NewMetaProperty(key types.ID, data types.Data) types.MetaProperty {
	return metaProperty{
		ID:   baseIDs.NewPropertyID(key, data.GetType()),
		Data: data,
	}
}

func ReadMetaProperty(metaPropertyString string) (types.MetaProperty, error) {
	propertyIDAndData := strings.Split(metaPropertyString, constants.PropertyIDAndDataSeparator)
	if len(propertyIDAndData) == 2 && propertyIDAndData[0] != "" {
		data, Error := utilities.ReadData(propertyIDAndData[1])
		if Error != nil {
			return nil, Error
		}

		return NewMetaProperty(baseIDs.NewID(propertyIDAndData[0]), data), nil
	}

	return nil, errors.IncorrectFormat
}
