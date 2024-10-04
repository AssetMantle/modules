package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Record interface {
	GetKey() Key
	GetMappable() Mappable

	WithKey(Key) Record

	ReadFromIterator(sdkTypes.Iterator) Record
	Read(sdkTypes.KVStore) Record
	Write(sdkTypes.KVStore) Record
	Delete(sdkTypes.KVStore)
}

func RecordsFromImplementations[T Record](records []T) []Record {
	Records := make([]Record, len(records))
	for i, record := range records {
		Records[i] = record
	}

	return Records
}
