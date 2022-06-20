// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type id struct {
	IDString string `json:"idString"`
}

var _ ids.ID = (*id)(nil)

func (id id) String() string {
	return id.IDString
}
func (id id) Bytes() []byte {
	return []byte(id.IDString)
}
func (id id) Compare(listable traits.Listable) int {
	if id, err := idFromInterface(listable); err != nil {
		panic(err)
	} else {
		return strings.Compare(id.String(), id.String())
	}
}
func idFromInterface(i interface{}) (id, error) {
	switch value := i.(type) {
	case id:
		return value, nil
	default:
		return id{}, constants.MetaDataError
	}
}

func NewID(idString string) ids.ID {
	return id{IDString: idString}
}
