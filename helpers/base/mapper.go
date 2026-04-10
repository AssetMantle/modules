// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"context"
	"fmt"
	prefixStore "cosmossdk.io/store/prefix"
	storeTypes "cosmossdk.io/store/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/AssetMantle/modules/helpers"
)

type mapper struct {
	kvStoreKey      *storeTypes.KVStoreKey
	recordPrototype func() helpers.Record
}

var _ helpers.Mapper = (*mapper)(nil)

func (mapper mapper) NewCollection(context context.Context) helpers.Collection {
	return collection{}.Initialize(context, mapper)
}
func (mapper mapper) GetKVStoreKey() *storeTypes.KVStoreKey {
	return mapper.kvStoreKey
}
func (mapper mapper) Upsert(context context.Context, record helpers.Record) {
	record.Write(sdkTypes.UnwrapSDKContext(context).KVStore(mapper.kvStoreKey))
}
func (mapper mapper) Read(context context.Context, key helpers.Key) helpers.Record {
	return mapper.recordPrototype().WithKey(key).Read(sdkTypes.UnwrapSDKContext(context).KVStore(mapper.kvStoreKey))
}
func (mapper mapper) Delete(context context.Context, key helpers.Key) {
	mapper.recordPrototype().WithKey(key).Delete(sdkTypes.UnwrapSDKContext(context).KVStore(mapper.kvStoreKey))
}
func (mapper mapper) FetchAll(context context.Context) []helpers.Record {
	var records []helpers.Record
	mapper.IterateAll(context, func(record helpers.Record) bool {
		records = append(records, record)
		return false
	})
	return records
}
func (mapper mapper) Iterate(context context.Context, key helpers.Key, accumulator func(helpers.Record) bool) {
	store := sdkTypes.UnwrapSDKContext(context).KVStore(mapper.kvStoreKey)
	prefixBytes := key.GenerateStorePrefixBytes()
	keyBytes := key.GenerateStoreKeyBytes()

	// SDK v0.50 panics on empty iterator keys in KVStorePrefixIterator.
	// When key bytes are empty (IterateAll / partial key), use Iterator(nil, nil)
	// on the prefix store to iterate all entries under that prefix.
	// When both prefix and key bytes are empty (full prototype), iterate the raw
	// store but skip the parameters store prefix (0x01) which shares the same
	// KV store.
	var iter storeTypes.Iterator
	skipParams := false
	if len(prefixBytes) == 0 && len(keyBytes) == 0 {
		iter = store.Iterator(nil, nil)
		skipParams = true
	} else if len(keyBytes) == 0 {
		iter = prefixStore.NewStore(store, prefixBytes).Iterator(nil, nil)
	} else {
		iter = storeTypes.KVStorePrefixIterator(prefixStore.NewStore(store, prefixBytes), keyBytes)
	}

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		if skipParams && bytes.Equal(iter.Key(), parametersStorePrefix) {
			continue
		}
		if accumulator(mapper.recordPrototype().ReadFromIterator(iter)) {
			break
		}
	}
}
func (mapper mapper) IteratePaginated(context context.Context, key helpers.Key, limit int32, accumulator func(helpers.Record) bool) {
	mapper.Iterate(context, key, func(record helpers.Record) bool {
		if limit > 0 {
			limit--
			return accumulator(record)
		}
		return true
	})
}
func (mapper mapper) IterateAll(context context.Context, accumulator func(record helpers.Record) bool) {
	mapper.Iterate(context, mapper.recordPrototype().GetKey(), accumulator)
}
func (mapper mapper) StoreDecoder(kvA kv.Pair, kvB kv.Pair) string {
	if bytes.Equal(kvA.Key[:1], mapper.recordPrototype().GetKey().GeneratePrefixedStoreKeyBytes()) {
		mappableA := mapper.recordPrototype().GetMappable()
		CodecPrototype().MustUnmarshal(kvA.Value, mappableA)

		mappableB := mapper.recordPrototype().GetMappable()
		CodecPrototype().MustUnmarshal(kvB.Value, mappableB)

		return fmt.Sprintf("%v\n%v", mappableA, mappableB)
	}

	panic(fmt.Errorf("invalid key prefix %X", kvA.Key[:1]))
}
func (mapper mapper) Initialize(kvStoreKey *storeTypes.KVStoreKey) helpers.Mapper {
	mapper.kvStoreKey = kvStoreKey
	return mapper
}
func NewMapper(recordPrototype func() helpers.Record) helpers.Mapper {
	return mapper{
		recordPrototype: recordPrototype,
	}
}
