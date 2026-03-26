package helpers

import (
	storeTypes "cosmossdk.io/store/types"
)

type Record interface {
	GetKey() Key
	GetMappable() Mappable

	WithKey(Key) Record

	ReadFromIterator(storeTypes.Iterator) Record
	Read(storeTypes.KVStore) Record
	Write(storeTypes.KVStore) Record
	Delete(storeTypes.KVStore)
}

func RecordsFromImplementations[T Record](records []T) []Record {
	Records := make([]Record, len(records))
	for i, record := range records {
		Records[i] = record
	}

	return Records
}

func RecordsToImplementations[T Record](_ T, records []Record) []T {
	implementations := make([]T, len(records))
	for i, record := range records {
		implementations[i] = record.(T)
	}
	return implementations
}
