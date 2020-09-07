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
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/tendermint/tendermint/libs/kv"
)

type mapper struct {
	kvStoreKey        *sdkTypes.KVStoreKey
	parameters        helpers.Parameters
	keyGenerator      func(types.ID) []byte
	mappablePrototype func() traits.Mappable
	registerCodec     func(*codec.Codec)
}

var _ helpers.Mapper = (*mapper)(nil)

func (mapper mapper) GetKVStoreKey() *sdkTypes.KVStoreKey {
	return mapper.kvStoreKey
}
func (mapper mapper) Create(context sdkTypes.Context, mappable traits.Mappable) {
	Bytes := mappable.Encode()
	kvStore := context.KVStore(mapper.kvStoreKey)
	kvStore.Set(mapper.keyGenerator(mappable.GetID()), Bytes)
}
func (mapper mapper) Read(context sdkTypes.Context, id types.ID) traits.Mappable {
	kvStore := context.KVStore(mapper.kvStoreKey)
	Bytes := kvStore.Get(mapper.keyGenerator(id))
	if Bytes == nil {
		return nil
	}
	return mapper.mappablePrototype().Decode(Bytes)
}
func (mapper mapper) Update(context sdkTypes.Context, mappable traits.Mappable) {
	Bytes := mappable.Encode()
	id := mappable.GetID()
	kvStore := context.KVStore(mapper.kvStoreKey)
	kvStore.Set(mapper.keyGenerator(id), Bytes)
}
func (mapper mapper) Delete(context sdkTypes.Context, id types.ID) {
	kvStore := context.KVStore(mapper.kvStoreKey)
	kvStore.Delete(mapper.keyGenerator(id))
}
func (mapper mapper) Iterate(context sdkTypes.Context, id types.ID, accumulator func(traits.Mappable) bool) {
	store := context.KVStore(mapper.kvStoreKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, mapper.keyGenerator(id))

	defer kvStorePrefixIterator.Close()
	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		mappable := mapper.mappablePrototype().Decode(kvStorePrefixIterator.Value())
		if accumulator(mappable) {
			break
		}
	}
}
func (mapper mapper) StoreDecoder(_ *codec.Codec, kvA kv.Pair, kvB kv.Pair) string {
	if bytes.Equal(kvA.Key[:1], mapper.keyGenerator(base.NewID(""))) {
		return fmt.Sprintf("%v\n%v", mapper.mappablePrototype().Decode(kvA.Value), mapper.mappablePrototype().Decode(kvB.Value))
	} else {
		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}
func (mapper mapper) RegisterCodec(codec *codec.Codec) {
	mapper.registerCodec(codec)
}

func NewMapper(module string, parameters helpers.Parameters, keyGenerator func(types.ID) []byte, mappablePrototype func() traits.Mappable, registerCodec func(*codec.Codec)) helpers.Mapper {
	return mapper{
		kvStoreKey:        sdkTypes.NewKVStoreKey(module),
		parameters:        parameters,
		keyGenerator:      keyGenerator,
		mappablePrototype: mappablePrototype,
		registerCodec:     registerCodec,
	}
}
