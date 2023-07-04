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

	Iterate(Key, func(Mappable) bool)
	IteratePaginated(Key, int32, func(Record) bool)
	Fetch(Key) Collection
	FetchRecord(Key) Record
	FetchAll() Collection
	FetchPaginated(Key, int32) Collection
	Add(Mappable) Collection
	Remove(Mappable) Collection
	Mutate(Mappable) Collection
	Initialize(context.Context, Mapper) Collection
}
