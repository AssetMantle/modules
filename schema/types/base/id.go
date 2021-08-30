/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strings"

	"github.com/persistenceOne/persistenceSDK/schema/types"
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
func (id id) Compare(compareID types.ID) int {
	return strings.Compare(id.String(), compareID.String())
}

func NewID(idString string) types.ID {
	return id{IDString: idString}
}
