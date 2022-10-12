// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/binary"
	"strconv"

	"github.com/AssetMantle/modules/schema/types"
)

type height struct {
	Value int64 `json:"height"`
}

func (height height) String() string {
	return strconv.FormatInt(height.Get(), 10)
}

func (height height) Bytes() []byte {
	Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(Bytes, uint64(height.Get()))
	return Bytes
}

var _ types.Height = (*height)(nil)

func (height height) Get() int64 { return height.Value }
func (height height) Compare(compareHeight types.Height) int {
	if height.Get() > compareHeight.Get() {
		return 1
	} else if height.Get() < compareHeight.Get() {
		return -1
	}

	return 0
}

func NewHeight(value int64) types.Height {
	if value < 0 {
		value = -1
	}

	return height{Value: value}
}
