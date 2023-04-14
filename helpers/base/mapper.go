// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"context"
	"fmt"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/AssetMantle/modules/helpers"
)

type mapper struct {
	kvStoreKey        *sdkTypes.KVStoreKey
	keyPrototype      func() helpers.Key
	mappablePrototype func() helpers.Mappable
}

var _ helpers.Mapper = (*mapper)(nil)

func (mapper mapper) NewCollection(context context.Context) helpers.Collection {
	return collection{}.Initialize(context, mapper)
}
func (mapper mapper) GetKVStoreKey() *sdkTypes.KVStoreKey {
	return mapper.kvStoreKey
}
func (mapper mapper) Create(context context.Context, mappable helpers.Mappable) {
	Bytes := CodecPrototype().MustMarshal(mappable)
	kvStore := sdkTypes.UnwrapSDKContext(context).KVStore(mapper.kvStoreKey)
	kvStore.Set(mappable.GetKey().GenerateStoreKeyBytes(), Bytes)
}
func (mapper mapper) Read(context context.Context, key helpers.Key) helpers.Mappable {
	kvStore := sdkTypes.UnwrapSDKContext(context).KVStore(mapper.kvStoreKey)

	Bytes := kvStore.Get(key.GenerateStoreKeyBytes())
	if Bytes == nil {
		return nil
	}

	mappable := mapper.mappablePrototype()
	CodecPrototype().MustUnmarshal(Bytes, mappable)

	return mappable
}
func (mapper mapper) Update(context context.Context, mappable helpers.Mappable) {
	Bytes := CodecPrototype().MustMarshal(mappable)
	key := mappable.GetKey()
	kvStore := sdkTypes.UnwrapSDKContext(context).KVStore(mapper.kvStoreKey)
	kvStore.Set(key.GenerateStoreKeyBytes(), Bytes)
}
func (mapper mapper) Delete(context context.Context, key helpers.Key) {
	kvStore := sdkTypes.UnwrapSDKContext(context).KVStore(mapper.kvStoreKey)
	kvStore.Delete(key.GenerateStoreKeyBytes())
}
func (mapper mapper) Iterate(context context.Context, partialKey helpers.Key, accumulator func(helpers.Mappable) bool) {
	store := sdkTypes.UnwrapSDKContext(context).KVStore(mapper.kvStoreKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, partialKey.GenerateStoreKeyBytes())

	defer func(kvStorePrefixIterator sdkTypes.Iterator) {
		err := kvStorePrefixIterator.Close()
		if err != nil {
			sdkTypes.UnwrapSDKContext(context).Logger().Debug(err.Error())
		}
	}(kvStorePrefixIterator)

	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		mappable := mapper.mappablePrototype()
		CodecPrototype().MustUnmarshal(kvStorePrefixIterator.Value(), mappable)
		if accumulator(mappable) {
			break
		}
	}
}
func (mapper mapper) IterateAll(context context.Context, accumulator func(helpers.Mappable) bool) {
	mapper.Iterate(context, mapper.keyPrototype(), accumulator)
}
func (mapper mapper) ReverseIterate(context context.Context, partialKey helpers.Key, accumulator func(helpers.Mappable) bool) {
	store := sdkTypes.UnwrapSDKContext(context).KVStore(mapper.kvStoreKey)
	kvStoreReversePrefixIterator := sdkTypes.KVStoreReversePrefixIterator(store, partialKey.GenerateStoreKeyBytes())

	defer func(kvStoreReversePrefixIterator sdkTypes.Iterator) {
		err := kvStoreReversePrefixIterator.Close()
		if err != nil {
			sdkTypes.UnwrapSDKContext(context).Logger().Debug(err.Error())
		}
	}(kvStoreReversePrefixIterator)

	for ; kvStoreReversePrefixIterator.Valid(); kvStoreReversePrefixIterator.Next() {
		mappable := mapper.mappablePrototype()
		CodecPrototype().MustUnmarshal(kvStoreReversePrefixIterator.Value(), mappable)

		if accumulator(mappable) {
			break
		}
	}
}
func (mapper mapper) StoreDecoder(kvA kv.Pair, kvB kv.Pair) string {
	if bytes.Equal(kvA.Key[:1], mapper.keyPrototype().GenerateStoreKeyBytes()) {
		mappableA := mapper.mappablePrototype()
		CodecPrototype().MustUnmarshal(kvA.Value, mappableA)

		mappableB := mapper.mappablePrototype()
		CodecPrototype().MustUnmarshal(kvB.Value, mappableB)

		return fmt.Sprintf("%v\n%v", mappableA, mappableB)
	}

	panic(fmt.Errorf("invalid key prefix %X", kvA.Key[:1]))
}
func (mapper mapper) Initialize(kvStoreKey *sdkTypes.KVStoreKey) helpers.Mapper {
	mapper.kvStoreKey = kvStoreKey
	return mapper
}
func NewMapper(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) helpers.Mapper {
	return mapper{
		keyPrototype:      keyPrototype,
		mappablePrototype: mappablePrototype,
	}
}
