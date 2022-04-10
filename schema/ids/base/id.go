// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/traits"
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
func (id id) Compare(listable traits.Listable) int {
	return strings.Compare(id.String(), idFromInterface(listable).String())
}
func idFromInterface(i interface{}) id {
	// TODO implement
	panic("")
}

func NewID(idString string) types.ID {
	return id{IDString: idString}
}
