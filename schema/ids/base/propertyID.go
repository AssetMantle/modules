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

//type propertyID struct {
//	Key  ids.StringID
//	Type ids.StringID
//}

var _ ids.PropertyID = (*PropertyID)(nil)

func (propertyID *PropertyID) IsPropertyID() {}
func (propertyID *PropertyID) GetKey() ids.StringID {
	return propertyID.KeyId
}
func (propertyID *PropertyID) GetType() ids.StringID {
	return propertyID.TypeId
}
func (propertyID *PropertyID) PropertyIDString() string {
	return stringUtilities.JoinIDStrings(propertyID.KeyId.String(), propertyID.TypeId.String())
}
func (propertyID *PropertyID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, propertyID.KeyId.Bytes()...)
	Bytes = append(Bytes, propertyID.TypeId.Bytes()...)

	return Bytes
}
func (propertyID *PropertyID) Compare(listable traits.Listable) int {
	return bytes.Compare(propertyID.Bytes(), propertyIDFromInterface(listable).Bytes())
}
func (propertyID *PropertyID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_PropertyId{
			PropertyId: propertyID,
		},
	}
}

func propertyIDFromInterface(listable traits.Listable) *PropertyID {
	switch value := listable.(type) {
	case *PropertyID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewPropertyID(key, Type ids.StringID) ids.PropertyID {
	return &PropertyID{
		KeyId:  key.(*StringID),
		TypeId: Type.(*StringID),
	}
}
