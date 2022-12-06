// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// type propertyID struct {
//	Key  ids.StringID
//	Type ids.StringID
// }

type propertyID base.PropertyID

func (propertyID *propertyID) String() string {
	// TODO implement me
	panic("implement me")
}

var _ ids.PropertyID = (*propertyID)(nil)

func (propertyID *propertyID) IsPropertyID() {}
func (propertyID *propertyID) GetKey() ids.StringID {
	return propertyID.KeyID
}
func (propertyID *propertyID) GetType() ids.StringID {
	return propertyID.TypeID
}
func (propertyID *propertyID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, propertyID.KeyID.Bytes()...)
	Bytes = append(Bytes, propertyID.TypeID.Bytes()...)

	return Bytes
}
func (propertyID *propertyID) Compare(listable traits.Listable) int {
	return bytes.Compare(propertyID.Bytes(), propertyIDFromInterface(listable).Bytes())
}
func propertyIDFromInterface(listable traits.Listable) *propertyID {
	switch value := listable.(type) {
	case *propertyID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewPropertyID(key, Type ids.StringID) ids.PropertyID {
	return &propertyID{
		KeyID:  key.(*stringID),
		TypeID: Type.(*stringID),
	}
}
