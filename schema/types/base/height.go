// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/binary"
	"strconv"

	"github.com/AssetMantle/modules/schema/types"
)

var _ types.Height = (*Height)(nil)

func (height *Height) StringHeight() string {
	return strconv.FormatInt(height.Get(), 10)
}

func (height *Height) Bytes() []byte {
	Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(Bytes, uint64(height.Get()))
	return Bytes
}

func (height *Height) Get() int64 { return height.Value }
func (height *Height) Compare(compareHeight types.Height) int {
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

	return &Height{Value: value}
}
