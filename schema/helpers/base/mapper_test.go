// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/utilities/test"
	"github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

func TestMapper(t *testing.T) {
	context, storeKey, _ := test.SetupTest(t)

	// NewMapper
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	testMapper := Mapper.(mapper)

	// Initialize
	require.Equal(t, storeKey, testMapper.kvStoreKey)

	// GetKVStoreKey
	require.Equal(t, "test", testMapper.GetKVStoreKey().Name())

	// NewCollection
	require.Equal(t, reflect.TypeOf(collection{}), reflect.TypeOf(testMapper.NewCollection(context).(collection)))
	require.Equal(t, testMapper.kvStoreKey.String(), testMapper.NewCollection(context).(collection).mapper.(mapper).kvStoreKey.String())
	require.Equal(t, context, testMapper.NewCollection(context).(collection).context)

	// Create
	testMapper.Create(context, base.NewMappable("test1", "value1"))
	testMapper.Create(context, base.NewMappable("test2", "value2"))
	testMapper.Create(context, base.NewMappable("test3", "value3"))

	// Set
	testMapper.Update(context, base.NewMappable("test2", "value3"))

	// Remove
	testMapper.Delete(context, base.NewKey("test3"))

	// Read
	require.Equal(t, base.NewMappable("test1", "value1"), testMapper.Read(context, base.NewKey("test1")))
	require.Equal(t, base.NewMappable("test2", "value3"), testMapper.Read(context, base.NewKey("test2")))
	require.Equal(t, nil, testMapper.Read(context, base.NewKey("test3")))

	// Iterate
	testMapper.Iterate(context, base.NewKey("test1"), func(mappable helpers.Mappable) bool { return false })
	testMapper.Iterate(context, base.NewKey("test3"), func(mappable helpers.Mappable) bool { return false })

	// TODO: Test functionality without legacy amino
	// Store Decoder
	// require.Equal(t, "{test1 value1}\n{test1 value1}", testMapper.StoreDecoder(kv.Pair{
	//	Key: append([]byte{0x11}, []byte("test1")...), Value: testMapper..MustMarshalBinaryBare(base.NewMappable("test1", "value1"))}, kv.Pair{
	//	Key: append([]byte{0x11}, []byte("test1")...), Value: testMapper.legacyAmino.MustMarshalBinaryBare(base.NewMappable("test1", "value1"))}),
	// )

}
