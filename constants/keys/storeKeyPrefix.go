// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package keys

import "encoding/binary"

// TODO move to helpers
type storeKeyPrefix int16

// TODO use enum
const (
	Assets storeKeyPrefix = iota + 8
	Classifications
	Identities
	Maintainers
	Metas
	Orders
	Splits
)

// TODO migrate to utilities
func (storeKeyPrefix storeKeyPrefix) GenerateStoreKey(key []byte) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(storeKeyPrefix))

	return append(bytes, key...)
}
