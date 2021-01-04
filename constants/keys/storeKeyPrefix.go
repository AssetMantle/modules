package keys

import "encoding/binary"

type storeKeyPrefix int16

const (
	Assets storeKeyPrefix = iota + 8
	Classifications
	Identities
	Maintainers
	Metas
	Orders
	Splits
)

func (storeKeyPrefix storeKeyPrefix) GenerateStoreKey(key []byte) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(storeKeyPrefix))
	return append(bytes, key...)
}
