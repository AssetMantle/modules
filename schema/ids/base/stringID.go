// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.StringID = (*StringID)(nil)

func (stringID *StringID) IsStringID() {}
func (stringID *StringID) AsString() string {
	return stringID.IDString
}
func (stringID *StringID) Bytes() []byte {
	return []byte(stringID.IDString)
}
func (stringID *StringID) Compare(listable traits.Listable) int {
	return strings.Compare(stringID.AsString(), stringIDFromInterface(listable).AsString())
}
func (stringID *StringID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_StringID{
			StringID: stringID,
		},
	}
}

func stringIDFromInterface(i interface{}) *StringID {
	switch value := i.(type) {
	case *StringID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewStringID(idString string) ids.StringID {
	return &StringID{IDString: idString}
}

func PrototypeStringID() ids.StringID {
	return &StringID{
		IDString: "",
	}
}
