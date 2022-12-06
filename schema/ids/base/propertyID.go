// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// type propertyID struct {
//	Key  ids.StringID
//	Type ids.StringID
// }

var _ ids.PropertyID = (*PropertyIDI_PropertyID)(nil)

func (propertyID *PropertyIDI_PropertyID) String() string {
	return propertyID.PropertyID.String()
}
func (propertyID *PropertyIDI_PropertyID) IsPropertyID() {}
func (propertyID *PropertyIDI_PropertyID) GetKey() ids.StringID {
	return propertyID.PropertyID.KeyID
}
func (propertyID *PropertyIDI_PropertyID) GetType() ids.StringID {
	return propertyID.PropertyID.TypeID
}
func (propertyID *PropertyIDI_PropertyID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, propertyID.PropertyID.KeyID.Bytes()...)
	Bytes = append(Bytes, propertyID.PropertyID.TypeID.Bytes()...)

	return Bytes
}
func (propertyID *PropertyIDI_PropertyID) Compare(listable traits.Listable) int {
	return bytes.Compare(propertyID.Bytes(), propertyIDFromInterface(listable).Bytes())
}
func propertyIDFromInterface(listable traits.Listable) *PropertyIDI_PropertyID {
	switch value := listable.(type) {
	case *PropertyIDI_PropertyID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func GeneratePropertyID(key, Type ids.StringID) ids.PropertyID {
	return NewPropertyID(key, Type)
}
func NewPropertyID(keyID, typeID ids.StringID) ids.PropertyID {
	return &PropertyIDI{
		Impl: &PropertyIDI_PropertyID{
			PropertyID: &PropertyID{
				KeyID:  keyID.(*StringIDI),
				TypeID: typeID.(*StringIDI),
			},
		},
	}

}
