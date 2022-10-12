package base

import (
	"encoding/binary"

	"github.com/AssetMantle/modules/schema/helpers"
)

type storeKeyPrefix int8

var _ helpers.StoreKeyPrefix = (*storeKeyPrefix)(nil)

func (storeKeyPrefix storeKeyPrefix) GenerateStoreKey(key []byte) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(storeKeyPrefix))

	return append(bytes, key...)
}

func NewStoreKeyPrefix(value int8) helpers.StoreKeyPrefix {
	return storeKeyPrefix(value)
}
