/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import "github.com/persistenceOne/persistenceSDK/schema/types"

var _ types.Height = (*Height)(nil)

func (height Height) Get() int64 { return height.Value }
func (height Height) Compare(compareHeight types.Height) int {
	if height.Get() > compareHeight.Get() {
		return 1
	} else if height.Get() < compareHeight.Get() {
		return -1
	}

	return 0
}
func NewHeight(value int64) *Height {
	return &Height{Value: value}
}
