/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.ID = (*ID)(nil)

func (id ID) String() string {
	return id.IdString
}
func (id ID) Bytes() []byte {
	return []byte(id.IdString)
}
func (id ID) Compare(compareID types.ID) int {
	return bytes.Compare(id.Bytes(), compareID.Bytes())
}

func NewID(idString string) types.ID {
	return ID{IdString: idString}
}
