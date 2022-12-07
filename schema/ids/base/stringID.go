// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.StringID = (*ID_StringID)(nil)

func (stringID *ID_StringID) String() string {
	return stringID.StringID.IdString
}
func (stringID *ID_StringID) IsStringID() {}
func (stringID *ID_StringID) Bytes() []byte {
	return []byte(stringID.String())
}
func (stringID *ID_StringID) Compare(listable traits.Listable) int {
	return strings.Compare(stringID.String(), idFromInterface(listable).String())
}

func NewStringID(idString string) ids.ID {
	return &ID{
		Impl: &ID_StringID{
			StringID: &StringID{
				IdString: idString,
			},
		},
	}
}

func PrototypeStringID() ids.ID {
	return NewStringID("")
}
