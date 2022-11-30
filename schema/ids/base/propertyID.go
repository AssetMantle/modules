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
//type propertyID struct {
//	Key  ids.StringID
//	Type ids.StringID
//}

var _ ids.PropertyID = (*PropertyID)(nil)

func (propertyID PropertyID) IsPropertyID() {}

//func (propertyID PropertyID) GetType() ids.StringID {
//	return propertyID.Type
//}
//func (propertyID PropertyID) String() string {
//	return stringUtilities.JoinIDStrings(propertyID.Key.String(), propertyID.Type.String())
//}
func (propertyID PropertyID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, *propertyID.Key)
	Bytes = append(Bytes, *propertyID.Type)

	return Bytes
}
func (propertyID PropertyID) Compare(listable traits.Listable) int {
	return bytes.Compare(propertyID.Bytes(), propertyIDFromInterface(listable).Bytes())
}
func propertyIDFromInterface(listable traits.Listable) PropertyID {
	switch value := listable.(type) {
	case PropertyID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewPropertyID(key, Type ids.StringID) ids.PropertyID {
	return &PropertyID{
		Key:  key.(*StringID),
		Type: Type.(*StringID),
	}
}
