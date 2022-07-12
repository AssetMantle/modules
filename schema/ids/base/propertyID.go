// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

type propertyID struct {
	Key  ids.StringID
	Type ids.StringID
}

var _ ids.PropertyID = (*propertyID)(nil)

func (propertyID propertyID) GetKey() ids.ID {
	return propertyID.Key
}
func (propertyID propertyID) GetType() ids.ID {
	return propertyID.Type
}
func (propertyID propertyID) String() string {
	return stringUtilities.JoinIDStrings(propertyID.Key.String(), propertyID.Type.String())
}
func (propertyID propertyID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, propertyID.Key.Bytes()...)
	Bytes = append(Bytes, propertyID.Type.Bytes()...)

	return Bytes
}
func (propertyID propertyID) Compare(listable traits.Listable) int {
	if comparePropertyID, err := propertyIDFromInterface(listable); err != nil {
		panic(err)
	} else {
		return bytes.Compare(propertyID.Bytes(), comparePropertyID.Bytes())

	}
}
func propertyIDFromInterface(listable traits.Listable) (propertyID, error) {
	switch value := listable.(type) {
	case propertyID:
		return value, nil
	default:
		return propertyID{}, errorConstants.MetaDataError
	}
}

func NewPropertyID(key, Type ids.StringID) ids.PropertyID {
	return propertyID{
		Key:  key,
		Type: Type,
	}
}
