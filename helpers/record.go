package helpers

import "github.com/cosmos/cosmos-sdk/types"

type Record interface {
	GetKey() Key
	GetMappable() Mappable

	WithKey(Key) Record
	WithMappable(Mappable) Record

	ReadFromIterator(types.Iterator) Record
	Read(types.KVStore) Record
	Write(types.KVStore) Record
	Delete(types.KVStore)
}
