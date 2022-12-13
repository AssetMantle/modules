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

func (propertyID propertyID) IsPropertyID() {}
func (propertyID propertyID) GetKey() ids.StringID {
	return propertyID.Key
}
func (propertyID propertyID) GetType() ids.StringID {
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
	return bytes.Compare(propertyID.Bytes(), propertyIDFromInterface(listable).Bytes())
}
func propertyIDFromInterface(listable traits.Listable) propertyID {
	switch value := listable.(type) {
	case propertyID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewPropertyID(key, Type ids.StringID) ids.PropertyID {
	return propertyID{
		Key:  key,
		Type: Type,
	}
}
