// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"fmt"

	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
)

type mapper struct {
	kvStoreKey        *sdkTypes.KVStoreKey
	codec             helpers.Codec
	keyPrototype      func() helpers.Key
	mappablePrototype func() helpers.Mappable
}

var _ helpers.Mapper = (*mapper)(nil)

func (mapper mapper) NewCollection(context sdkTypes.Context) helpers.Collection {
	return collection{}.Initialize(context, mapper)
}

func (mapper mapper) GetKVStoreKey() *sdkTypes.KVStoreKey {
	return mapper.kvStoreKey
}
func (mapper mapper) Create(context sdkTypes.Context, mappable helpers.Mappable) {
	Bytes := mapper.codec.MustMarshal(mappable)
	kvStore := context.KVStore(mapper.kvStoreKey)
	kvStore.Set(mappable.GetKey().GenerateStoreKeyBytes(), Bytes)
}
func (mapper mapper) Read(context sdkTypes.Context, key helpers.Key) helpers.Mappable {
	kvStore := context.KVStore(mapper.kvStoreKey)

	Bytes := kvStore.Get(key.GenerateStoreKeyBytes())
	if Bytes == nil {
		return nil
	}

	var mappable helpers.Mappable

	mapper.codec.MustUnmarshal(Bytes, mappable)

	return mappable
}
func (mapper mapper) Update(context sdkTypes.Context, mappable helpers.Mappable) {
	Bytes := mapper.codec.MustMarshal(mappable)
	key := mappable.GetKey()
	kvStore := context.KVStore(mapper.kvStoreKey)
	kvStore.Set(key.GenerateStoreKeyBytes(), Bytes)
}
func (mapper mapper) Delete(context sdkTypes.Context, key helpers.Key) {
	kvStore := context.KVStore(mapper.kvStoreKey)
	kvStore.Delete(key.GenerateStoreKeyBytes())
}
func (mapper mapper) Iterate(context sdkTypes.Context, partialKey helpers.Key, accumulator func(helpers.Mappable) bool) {
	store := context.KVStore(mapper.kvStoreKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, partialKey.GenerateStoreKeyBytes())

	defer func(kvStorePrefixIterator sdkTypes.Iterator) {
		err := kvStorePrefixIterator.Close()
		if err != nil {
			context.Logger().Debug(err.Error())
		}
	}(kvStorePrefixIterator)

	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		var mappable helpers.Mappable
		mapper.codec.MustUnmarshal(kvStorePrefixIterator.Value(), mappable)
		if accumulator(mappable) {
			break
		}
	}
}
func (mapper mapper) ReverseIterate(context sdkTypes.Context, partialKey helpers.Key, accumulator func(helpers.Mappable) bool) {
	store := context.KVStore(mapper.kvStoreKey)
	kvStoreReversePrefixIterator := sdkTypes.KVStoreReversePrefixIterator(store, partialKey.GenerateStoreKeyBytes())

	defer func(kvStoreReversePrefixIterator sdkTypes.Iterator) {
		err := kvStoreReversePrefixIterator.Close()
		if err != nil {
			context.Logger().Debug(err.Error())
		}
	}(kvStoreReversePrefixIterator)

	for ; kvStoreReversePrefixIterator.Valid(); kvStoreReversePrefixIterator.Next() {
		var mappable helpers.Mappable

		mapper.codec.MustUnmarshal(kvStoreReversePrefixIterator.Value(), mappable)

		if accumulator(mappable) {
			break
		}
	}
}
func (mapper mapper) StoreDecoder(kvA kv.Pair, kvB kv.Pair) string {
	if bytes.Equal(kvA.Key[:1], mapper.keyPrototype().GenerateStoreKeyBytes()) {
		var mappableA helpers.Mappable

		mapper.codec.MustUnmarshal(kvA.Value, mappableA)

		var mappableB helpers.Mappable

		mapper.codec.MustUnmarshal(kvB.Value, mappableB)

		return fmt.Sprintf("%v\n%v", mappableA, mappableB)
	}

	panic(fmt.Errorf("invalid key prefix %X", kvA.Key[:1]))
}
func (mapper mapper) Initialize(kvStoreKey *sdkTypes.KVStoreKey) helpers.Mapper {
	mapper.kvStoreKey = kvStoreKey
	return mapper
}
func NewMapper(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) helpers.Mapper {
	legacyAmino := sdkCodec.NewLegacyAmino()
	keyPrototype().RegisterLegacyAminoCodec(legacyAmino)
	mappablePrototype().RegisterLegacyAminoCodec(legacyAmino)
	schema.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	return mapper{
		codec:             CodecPrototype().InitializeAndSeal(),
		keyPrototype:      keyPrototype,
		mappablePrototype: mappablePrototype,
	}
}
