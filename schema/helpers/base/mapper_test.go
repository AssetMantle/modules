/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/kv"
	"reflect"
	"testing"
)

func TestMapper(t *testing.T) {

	context, storeKey, _ := base.SetupTest(t)

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

	// Update
	testMapper.Update(context, base.NewMappable("test2", "value3"))

	// Delete
	testMapper.Delete(context, base.NewKey("test3"))

	// Read
	require.Equal(t, base.NewMappable("test1", "value1"), testMapper.Read(context, base.NewKey("test1")))
	require.Equal(t, base.NewMappable("test2", "value3"), testMapper.Read(context, base.NewKey("test2")))
	require.Equal(t, nil, testMapper.Read(context, base.NewKey("test3")))

	// Iterate
	testMapper.Iterate(context, base.NewKey("test1"), func(mappable helpers.Mappable) bool { return false })
	testMapper.Iterate(context, base.NewKey("test3"), func(mappable helpers.Mappable) bool { return false })

	// Store Decoder
	require.Equal(t, "{test1 value1}\n{test1 value1}", testMapper.StoreDecoder(codec.New(), kv.Pair{
		Key: append([]byte{0x11}, []byte("test1")...), Value: testMapper.codec.MustMarshalBinaryBare(base.NewMappable("test1", "value1"))}, kv.Pair{
		Key: append([]byte{0x11}, []byte("test1")...), Value: testMapper.codec.MustMarshalBinaryBare(base.NewMappable("test1", "value1"))}),
	)

}
