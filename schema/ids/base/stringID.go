// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.StringID = (*StringIDI_StringID)(nil)

func (stringID *StringIDI_StringID) String() string {
	return stringID.StringID.IdString
}
func (stringID *StringIDI_StringID) IsStringID() {}
func (stringID *StringIDI_StringID) Bytes() []byte {
	return []byte(stringID.String())
}
func (stringID *StringIDI_StringID) Compare(listable traits.Listable) int {
	return strings.Compare(stringID.String(), stringIDFromInterface(listable).Impl.(ids.StringID).String())
}
func stringIDFromInterface(i interface{}) *StringIDI {
	switch value := i.(type) {
	case *StringIDI:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewStringID(idString string) ids.StringID {
	return &StringIDI{
		Impl: &StringIDI_StringID{
			StringID: &StringID{
				IdString: idString,
			},
		},
	}
}

func PrototypeStringID() ids.StringID {
	return NewStringID("")
}
