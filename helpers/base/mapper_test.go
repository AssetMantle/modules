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

func Test_mapper_upsert_read_roundtrip(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)

	data := baseData.NewStringData("hello")
	rec := metasRecord.NewRecord(data)

	m.Upsert(ctx, rec)

	readBack := m.Read(ctx, rec.GetKey())
	require.NotNil(t, readBack)
	assert.Equal(t, rec.GetKey(), readBack.GetKey())

	mappable := readBack.GetMappable()
	require.NotNil(t, mappable)
	readData := metasMappable.GetData(mappable)
	require.NotNil(t, readData)
	assert.Contains(t, readData.AsString(), "hello")
}

func Test_mapper_delete_removes_record(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)

	data := baseData.NewStringData("toDelete")
	rec := metasRecord.NewRecord(data)
	m.Upsert(ctx, rec)

	m.Delete(ctx, rec.GetKey())

	readBack := m.Read(ctx, rec.GetKey())
	require.NotNil(t, readBack)
	// After delete, Read returns prototype with nil/empty mappable data
	readMappable := readBack.GetMappable()
	if readMappable != nil {
		readData := metasMappable.GetData(readMappable)
		assert.Nil(t, readData, "deleted record should return nil data")
	}
}

func Test_mapper_read_missing_returns_prototype(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)

	missingKey := metasKey.NewKey(baseIDs.GenerateDataID(baseData.NewStringData("nonexistent")))
	readBack := m.Read(ctx, missingKey)
	require.NotNil(t, readBack, "read of missing key should return prototype, not nil")
}

func Test_mapper_fetch_all(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)

	data1 := baseData.NewStringData("one")
	data2 := baseData.NewStringData("two")
	m.Upsert(ctx, metasRecord.NewRecord(data1))
	m.Upsert(ctx, metasRecord.NewRecord(data2))

	all := m.FetchAll(ctx)
	assert.Len(t, all, 2, "FetchAll should return all upserted records")
}

func Test_mapper_iterate_visits_all(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	m := baseHelpers.NewMapper(metasRecord.Prototype).Initialize(storeKey)

	data1 := baseData.NewStringData("alpha")
	data2 := baseData.NewStringData("beta")
	m.Upsert(ctx, metasRecord.NewRecord(data1))
	m.Upsert(ctx, metasRecord.NewRecord(data2))

	var visited int
	m.IterateAll(ctx, func(record helpers.Record) bool {
		visited++
		return false
	})
	assert.Equal(t, 2, visited, "IterateAll should visit all records")
}
