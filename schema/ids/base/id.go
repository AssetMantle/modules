// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/types"
)

type id struct {
	IDString string `json:"idString"`
}

var _ types.ID = (*id)(nil)

func (id id) String() string {
	return id.IDString
}
func (id id) Bytes() []byte {
	return []byte(id.IDString)
}
func (id id) Compare(listable capabilities.Listable) int {
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
		return id{}, errors.MetaDataError
	}
}

func NewID(idString string) types.ID {
	return id{IDString: idString}
}
