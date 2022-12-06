// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type stringID base.StringID

func (stringID *stringID) String() string {
	// TODO implement me
	panic("implement me")
}

var _ ids.StringID = (*stringID)(nil)

func (stringID *stringID) IsStringID() {}

func (stringID *stringID) Bytes() []byte {
	return []byte(stringID.IdString)
}
func (stringID *stringID) Compare(listable traits.Listable) int {
	return strings.Compare(stringID.String(), stringIDFromInterface(listable).String())
}
func stringIDFromInterface(i interface{}) *stringID {
	switch value := i.(type) {
	case *stringID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewStringID(idString string) ids.StringID {
	return &stringID{IdString: idString}
}

func PrototypeStringID() ids.StringID {
	return &stringID{
		IdString: "",
	}
}
