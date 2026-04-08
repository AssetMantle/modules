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
