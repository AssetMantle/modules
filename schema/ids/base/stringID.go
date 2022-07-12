// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type stringID struct {
	IDString string `json:"idString"`
}

var _ ids.StringID = (*stringID)(nil)

func (stringID stringID) String() string {
	return stringID.IDString
}
func (stringID stringID) Bytes() []byte {
	return []byte(stringID.IDString)
}
func (stringID stringID) Compare(listable traits.Listable) int {
	if id, err := stringIDFromInterface(listable); err != nil {
		panic(err)
	} else {
		return strings.Compare(id.String(), id.String())
	}
}
func stringIDFromInterface(i interface{}) (stringID, error) {
	switch value := i.(type) {
	case stringID:
		return value, nil
	default:
		return stringID{}, constants.MetaDataError
	}
}

func NewStringID(idString string) ids.StringID {
	return stringID{IDString: idString}
}
