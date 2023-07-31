// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
)

// Collection a list of mappable with create CRUD methods
type Collection interface {
	GetMappable(Key) Mappable
	Get() []Record
	GetMappables() []Mappable

	// TODO convert rest of the accumulators to below logic
	// The accumulator function should return true if the iterated Mappable is to be included in the returned collection
	IterateAll(func(Record) bool) Collection
	Iterate(Key, func(Record) bool)
	IteratePaginated(Key, int32, func(Record) bool)
	Fetch(Key) Collection
	FetchRecord(Key) Record
	FetchAll() Collection
	FetchPaginated(Key, int32) Collection
	Add(Record) Collection
	Remove(Record) Collection
	Mutate(Record) Collection
	Initialize(context.Context, Mapper) Collection
}
