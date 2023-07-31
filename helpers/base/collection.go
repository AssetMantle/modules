// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"

	"github.com/AssetMantle/modules/helpers"
)

type collection struct {
	records []helpers.Record

	mapper  helpers.Mapper
	context context.Context
}

var _ helpers.Collection = (*collection)(nil)

func (collection collection) GetMappable(key helpers.Key) helpers.Mappable {
	for _, record := range collection.records {
		if record.GetKey().Equals(key) {
			return record.GetMappable()
		}
	}

	return nil
}
func (collection collection) Get() []helpers.Record {
	return collection.records
}
func (collection collection) GetMappables() []helpers.Mappable {
	var mappables []helpers.Mappable
	for _, record := range collection.records {
		mappables = append(mappables, record.GetMappable())
	}
	return mappables
}
func (collection collection) IterateAll(accumulator func(record helpers.Record) bool) helpers.Collection {
	var records []helpers.Record

	collection.mapper.IterateAll(collection.context, func(record helpers.Record) bool {
		if accumulator(record) == true {
			records = append(records, record)
		}
		return false
	})

	collection.records = records

	return collection
}
func (collection collection) Iterate(startKey helpers.Key, accumulator func(helpers.Record) bool) {
	collection.mapper.Iterate(collection.context, startKey, accumulator)
}
func (collection collection) FetchAll() helpers.Collection {
	collection.records = collection.mapper.FetchAll(collection.context)
	return collection
}
func (collection collection) IteratePaginated(startKey helpers.Key, limit int32, accumulator func(helpers.Record) bool) {
	collection.mapper.IteratePaginated(collection.context, startKey, limit, accumulator)
}
func (collection collection) FetchRecord(key helpers.Key) helpers.Record {
	record := collection.mapper.Read(collection.context, key)
	collection.records = []helpers.Record{record}
	return record
}
func (collection collection) Fetch(key helpers.Key) helpers.Collection {
	var records []helpers.Record

	if key.IsPartial() {
		collection.mapper.Iterate(collection.context, key, func(record helpers.Record) bool {
			records = append(records, record)
			return false
		})
	} else {
		record := collection.mapper.Read(collection.context, key)
		if record != nil {
			records = append(records, record)
		}
	}

	collection.records = records

	return collection
}
func (collection collection) FetchPaginated(startKey helpers.Key, limit int32) helpers.Collection {
	var records []helpers.Record

	collection.mapper.IteratePaginated(collection.context, startKey, limit, func(record helpers.Record) bool {
		records = append(records, record)
		return false
	})

	collection.records = records

	return collection
}
func (collection collection) Add(record helpers.Record) helpers.Collection {
	collection.mapper.Upsert(collection.context, record)

	collection.records = []helpers.Record{record}

	return collection
}
func (collection collection) Remove(record helpers.Record) helpers.Collection {
	collection.mapper.Delete(collection.context, record.GetKey())

	collection.records = []helpers.Record{}

	return collection
}
func (collection collection) Mutate(record helpers.Record) helpers.Collection {
	collection.records = []helpers.Record{}

	if Record := collection.mapper.Read(collection.context, record.GetKey()); Record != nil {
		collection.mapper.Upsert(collection.context, record)
		collection.records = []helpers.Record{record}
	}

	return collection
}
func (collection collection) Initialize(context context.Context, mapper helpers.Mapper) helpers.Collection {
	collection.mapper = mapper
	collection.context = context

	return collection
}
