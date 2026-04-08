// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base_test

import (
	"testing"

	storeTypes "cosmossdk.io/store/types"
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	metasKey "github.com/AssetMantle/modules/x/metas/key"
	metasMappable "github.com/AssetMantle/modules/x/metas/mappable"
	metasRecord "github.com/AssetMantle/modules/x/metas/record"
)

func Test_collection_add_fetch_roundtrip(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data := baseData.NewStringData("testValue")
	rec := metasRecord.NewRecord(data)
	coll = coll.Add(rec)

	fetched := coll.Fetch(rec.GetKey())
	records := fetched.Get()
	require.Len(t, records, 1)
	assert.Equal(t, rec.GetKey(), records[0].GetKey())
}

func Test_collection_remove_deletes(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data := baseData.NewStringData("toRemove")
	rec := metasRecord.NewRecord(data)
	coll = coll.Add(rec)
	coll = coll.Remove(rec)

	// After removal, Get returns empty slice
	assert.Empty(t, coll.Get())

	// Re-fetch from store should return prototype (no data)
	fetched := coll.Fetch(rec.GetKey())
	records := fetched.Get()
	for _, r := range records {
		mappable := r.GetMappable()
		if mappable != nil {
			readData := metasMappable.GetData(mappable)
			assert.Nil(t, readData, "removed record should have nil data")
		}
	}
}

func Test_collection_mutate_updates_existing(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data := baseData.NewStringData("original")
	rec := metasRecord.NewRecord(data)
	coll = coll.Add(rec)

	coll = coll.Mutate(rec)
	records := coll.Get()
	require.Len(t, records, 1)
	assert.Equal(t, rec.GetKey(), records[0].GetKey())
}

func Test_collection_fetch_partial_key_iterates(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data1 := baseData.NewStringData("item1")
	data2 := baseData.NewStringData("item2")
	coll.Add(metasRecord.NewRecord(data1))
	coll.Add(metasRecord.NewRecord(data2))

	partialKey := metasKey.Prototype()
	require.True(t, partialKey.IsPartial())

	fetched := coll.Fetch(partialKey)
	records := fetched.Get()
	assert.GreaterOrEqual(t, len(records), 2, "partial key fetch should return multiple records")
}

func Test_collection_get_returns_records(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data := baseData.NewStringData("getValue")
	rec := metasRecord.NewRecord(data)
	coll = coll.Add(rec)

	got := coll.Get()
	require.Len(t, got, 1)
}

func Test_collection_get_mappable_nil_for_missing(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	missingKey := metasKey.NewKey(baseIDs.GenerateDataID(baseData.NewStringData("missing")))
	mappable := coll.GetMappable(missingKey)
	assert.Nil(t, mappable, "GetMappable for missing key should return nil")
}

func Test_collection_initialize(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)

	coll := m.NewCollection(ctx)
	require.NotNil(t, coll)
	assert.Empty(t, coll.Get())
}

func Test_collection_fetch_all(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data1 := baseData.NewStringData("all1")
	data2 := baseData.NewStringData("all2")
	coll.Add(metasRecord.NewRecord(data1))
	coll.Add(metasRecord.NewRecord(data2))

	fetchedAll := coll.FetchAll()
	records := fetchedAll.Get()
	assert.Len(t, records, 2)
}

func Test_collection_get_mappables(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data := baseData.NewStringData("mappableTest")
	rec := metasRecord.NewRecord(data)
	coll = coll.Add(rec)

	mappables := coll.GetMappables()
	require.Len(t, mappables, 1)

	var _ helpers.Mappable = mappables[0]
}

func Test_collection_iterate_all(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data1 := baseData.NewStringData("iterAll1")
	data2 := baseData.NewStringData("iterAll2")
	coll.Add(metasRecord.NewRecord(data1))
	coll.Add(metasRecord.NewRecord(data2))

	count := 0
	result := coll.IterateAll(func(record helpers.Record) bool {
		count++
		return true // accumulate
	})
	assert.Equal(t, 2, count)
	assert.Len(t, result.Get(), 2)
}

func Test_collection_iterate(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data1 := baseData.NewStringData("iter1")
	data2 := baseData.NewStringData("iter2")
	coll.Add(metasRecord.NewRecord(data1))
	coll.Add(metasRecord.NewRecord(data2))

	count := 0
	partialKey := metasKey.Prototype()
	coll.Iterate(partialKey, func(record helpers.Record) bool {
		count++
		return false
	})
	assert.GreaterOrEqual(t, count, 2)
}

func Test_collection_fetch_record(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data := baseData.NewStringData("fetchRecord")
	rec := metasRecord.NewRecord(data)
	coll.Add(rec)

	fetched := coll.FetchRecord(rec.GetKey())
	require.NotNil(t, fetched)
	assert.Equal(t, rec.GetKey(), fetched.GetKey())
}

func Test_collection_fetch_paginated(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data1 := baseData.NewStringData("page1")
	data2 := baseData.NewStringData("page2")
	data3 := baseData.NewStringData("page3")
	coll.Add(metasRecord.NewRecord(data1))
	coll.Add(metasRecord.NewRecord(data2))
	coll.Add(metasRecord.NewRecord(data3))

	partialKey := metasKey.Prototype()
	paginated := coll.FetchPaginated(partialKey, 2)
	records := paginated.Get()
	assert.LessOrEqual(t, len(records), 2)
	assert.Greater(t, len(records), 0)
}

func Test_collection_iterate_paginated(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)
	coll := m.NewCollection(ctx)

	data1 := baseData.NewStringData("iterPage1")
	data2 := baseData.NewStringData("iterPage2")
	data3 := baseData.NewStringData("iterPage3")
	coll.Add(metasRecord.NewRecord(data1))
	coll.Add(metasRecord.NewRecord(data2))
	coll.Add(metasRecord.NewRecord(data3))

	count := 0
	partialKey := metasKey.Prototype()
	coll.IteratePaginated(partialKey, 2, func(record helpers.Record) bool {
		count++
		return false
	})
	assert.LessOrEqual(t, count, 2)
	assert.Greater(t, count, 0)
}
