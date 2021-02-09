/*
Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strconv"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type heightID struct {
	Value int64 `json:"value"`
}

var _ types.ID = (*heightID)(nil)

func (h heightID) String() string {
	return strconv.FormatInt(h.Value, 10)
}

func (h heightID) Bytes() []byte {
	var Bytes []byte

	Bytes = append(Bytes, uint8(len(h.String())))
	Bytes = append(Bytes, []byte(h.String())...)

	return Bytes
}

func (h heightID) Equals(i types.ID) bool {
	switch v := i.(type) {
	case heightID:
		return h.Value == v.Value
	default:
		return false
	}
}

func NewHeightID(value int64) types.ID {
	return heightID{
		Value: value,
	}
}
