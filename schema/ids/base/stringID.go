// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

//type stringID struct {
//	IDString string `json:"idString"`
//}

var _ ids.StringID = (*StringID)(nil)

func (stringID *StringID) IsStringID() {}

func (stringID *StringID) Bytes() []byte {
	return []byte(stringID.IdString)
}
func (stringID *StringID) Compare(listable traits.Listable) int {
	return strings.Compare(stringID.String(), stringIDFromInterface(listable).String())
}
func (stringID *StringID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_StringId{
			StringId: stringID,
		},
	}
}

func stringIDFromInterface(i interface{}) *StringID {
	switch value := i.(type) {
	case *StringID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewStringID(idString string) ids.StringID {
	return &StringID{IdString: idString}
}

func PrototypeStringID() ids.StringID {
	return &StringID{
		IdString: "",
	}
}
