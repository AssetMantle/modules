/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/tendermint/tendermint/libs/kv"
)

type mapper struct {
	kvStoreKey        *sdkTypes.KVStoreKey
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
	Bytes := mappable.Encode()
	kvStore := context.KVStore(mapper.kvStoreKey)
	kvStore.Set(mappable.GetKey().GenerateStoreKeyBytes(), Bytes)
}
func (mapper mapper) Read(context sdkTypes.Context, key helpers.Key) helpers.Mappable {
	kvStore := context.KVStore(mapper.kvStoreKey)
	Bytes := kvStore.Get(key.GenerateStoreKeyBytes())
	if Bytes == nil {
		return nil
	}
	return mapper.mappablePrototype().Decode(Bytes)
}
func (mapper mapper) Update(context sdkTypes.Context, mappable helpers.Mappable) {
	Bytes := mappable.Encode()
	key := mappable.GetKey()
	kvStore := context.KVStore(mapper.kvStoreKey)
	kvStore.Set(key.GenerateStoreKeyBytes(), Bytes)
}
func (mapper mapper) Delete(context sdkTypes.Context, key helpers.Key) {
	kvStore := context.KVStore(mapper.kvStoreKey)
	kvStore.Delete(key.GenerateStoreKeyBytes())
}
func (mapper mapper) Iterate(context sdkTypes.Context, key helpers.Key, accumulator func(helpers.Mappable) bool) {
	store := context.KVStore(mapper.kvStoreKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, key.GenerateStoreKeyBytes())

	defer kvStorePrefixIterator.Close()
	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		mappable := mapper.mappablePrototype().Decode(kvStorePrefixIterator.Value())
		if accumulator(mappable) {
			break
		}
	}
}
func (mapper mapper) StoreDecoder(_ *codec.Codec, kvA kv.Pair, kvB kv.Pair) string {
	if bytes.Equal(kvA.Key[:1], mapper.keyPrototype().GenerateStoreKeyBytes()) {
		return fmt.Sprintf("%v\n%v", mapper.mappablePrototype().Decode(kvA.Value), mapper.mappablePrototype().Decode(kvB.Value))
	} else {
		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}
func (mapper mapper) RegisterCodec(codec *codec.Codec) {
	mapper.keyPrototype().RegisterCodec(codec)
	mapper.mappablePrototype().RegisterCodec(codec)
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
