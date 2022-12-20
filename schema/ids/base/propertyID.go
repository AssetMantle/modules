// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.PropertyID = (*ID_PropertyID)(nil)

func (propertyID *ID_PropertyID) String() string {
	return propertyID.PropertyID.String()
}
func (propertyID *ID_PropertyID) IsPropertyID() {}
func (propertyID *ID_PropertyID) GetKey() ids.ID {
	return &ID{Impl: &ID_StringID{StringID: propertyID.PropertyID.KeyID}}
}
func (propertyID *ID_PropertyID) GetType() ids.ID {
	return &ID{Impl: &ID_StringID{StringID: propertyID.PropertyID.TypeID}}
}
func (propertyID *ID_PropertyID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, propertyID.PropertyID.KeyID.IdString...)
	Bytes = append(Bytes, propertyID.PropertyID.TypeID.IdString...)

	return Bytes
}
func (propertyID *ID_PropertyID) Compare(listable traits.Listable) int {
	return bytes.Compare(propertyID.Bytes(), idFromInterface(listable).Bytes())
}

func GeneratePropertyID(key, Type ids.ID) ids.ID {
	return NewPropertyID(key, Type)
}
func NewPropertyID(keyID, typeID ids.ID) ids.ID {
	if keyID.(*ID).GetStringID() == nil || typeID.(*ID).GetStringID() == nil {
		panic(errorConstants.MetaDataError)
	}
	return &ID{
		Impl: &ID_PropertyID{
			PropertyID: &PropertyID{
				KeyID:  keyID.(*ID).GetStringID(),
				TypeID: typeID.(*ID).GetStringID(),
			},
		},
	}

}
