/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/kv"

	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type mapper struct {
	kvStoreKey        *sdkTypes.KVStoreKey
	legacyAminoCodec  *codec.LegacyAmino
	codec             codec.Marshaler
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
	Bytes := mapper.codec.MustMarshalBinaryBare(mappable)
	kvStore := context.KVStore(mapper.kvStoreKey)
	kvStore.Set(mappable.GetKey().GenerateStoreKeyBytes(), Bytes)
}
func (mapper mapper) Read(context sdkTypes.Context, key helpers.Key) helpers.Mappable {
	kvStore := context.KVStore(mapper.kvStoreKey)

	Bytes := kvStore.Get(key.GenerateStoreKeyBytes())
	if Bytes == nil {
		return nil
	}

	mappable := mapper.mappablePrototype().GetStructReference()

	mapper.codec.MustUnmarshalBinaryBare(Bytes, mappable)
	return mappable.(helpers.Mappable)
}
func (mapper mapper) Update(context sdkTypes.Context, mappable helpers.Mappable) {
	Bytes := mapper.codec.MustMarshalBinaryBare(mappable)
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

	defer kvStorePrefixIterator.Close()

	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		mappable := mapper.mappablePrototype().GetStructReference()

		mapper.codec.MustUnmarshalBinaryBare(kvStorePrefixIterator.Value(), mappable)

		if accumulator(mappable.(helpers.Mappable)) {
			break
		}
	}
}
func (mapper mapper) ReverseIterate(context sdkTypes.Context, partialKey helpers.Key, accumulator func(helpers.Mappable) bool) {
	store := context.KVStore(mapper.kvStoreKey)
	kvStoreReversePrefixIterator := sdkTypes.KVStoreReversePrefixIterator(store, partialKey.GenerateStoreKeyBytes())

	defer kvStoreReversePrefixIterator.Close()

	for ; kvStoreReversePrefixIterator.Valid(); kvStoreReversePrefixIterator.Next() {
		var mappable helpers.Mappable

		mapper.codec.MustUnmarshalBinaryBare(kvStoreReversePrefixIterator.Value(), mappable)

		if accumulator(mappable) {
			break
		}
	}
}
func (mapper mapper) StoreDecoder(kvA kv.Pair, kvB kv.Pair) string {
	if bytes.Equal(kvA.Key[:1], mapper.keyPrototype().GenerateStoreKeyBytes()) {
		var mappableA helpers.Mappable

		mapper.codec.MustUnmarshalBinaryBare(kvA.Value, mappableA)

		var mappableB helpers.Mappable

		mapper.codec.MustUnmarshalBinaryBare(kvB.Value, mappableB)

		return fmt.Sprintf("%v\n%v", mappableA, mappableB)
	}

	panic(fmt.Errorf("invalid key prefix %X", kvA.Key[:1]))
}
func (mapper mapper) Initialize(kvStoreKey *sdkTypes.KVStoreKey) helpers.Mapper {
	mapper.kvStoreKey = kvStoreKey
	return mapper
}
func NewMapper(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) helpers.Mapper {
	registry := codectypes.NewInterfaceRegistry()
	Codec := codec.NewProtoCodec(registry)
	LegacyAminoCodec := codec.NewLegacyAmino()
	keyPrototype().RegisterLegacyAminoCodec(LegacyAminoCodec)
	mappablePrototype().RegisterLegacyAminoCodec(LegacyAminoCodec)
	schema.RegisterLegacyAminoCodec(LegacyAminoCodec)
	LegacyAminoCodec.Seal()

	return mapper{
		codec:             Codec,
		legacyAminoCodec:  LegacyAminoCodec,
		keyPrototype:      keyPrototype,
		mappablePrototype: mappablePrototype,
	}
}
