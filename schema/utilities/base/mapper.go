package base

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type mapper struct {
	storeKey          sdkTypes.StoreKey
	codec             *codec.Codec
	keyGenerator      func(types.ID) []byte
	mappablePrototype func() traits.Mappable
	registerCodec     func(*codec.Codec)
}

var _ utilities.Mapper = (*mapper)(nil)

func (mapper mapper) Create(context sdkTypes.Context, mappable traits.Mappable) {
	bytes, Error := mapper.codec.MarshalBinaryBare(mappable)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.storeKey)
	kvStore.Set(mapper.keyGenerator(mappable.GetID()), bytes)
}
func (mapper mapper) Read(context sdkTypes.Context, id types.ID) traits.Mappable {
	kvStore := context.KVStore(mapper.storeKey)
	bytes := kvStore.Get(mapper.keyGenerator(id))
	if bytes == nil {
		return nil
	}
	mappable := mapper.mappablePrototype()
	Error := mapper.codec.UnmarshalBinaryBare(bytes, &mappable)
	if Error != nil {
		panic(Error)
	}
	return mappable
}
func (mapper mapper) Update(context sdkTypes.Context, mappable traits.Mappable) {
	bytes, Error := mapper.codec.MarshalBinaryBare(mappable)
	if Error != nil {
		panic(Error)
	}
	id := mappable.GetID()
	kvStore := context.KVStore(mapper.storeKey)
	kvStore.Set(mapper.keyGenerator(id), bytes)
}
func (mapper mapper) Delete(context sdkTypes.Context, id types.ID) {
	kvStore := context.KVStore(mapper.storeKey)
	kvStore.Delete(mapper.keyGenerator(id))
}
func (mapper mapper) Iterate(context sdkTypes.Context, id types.ID, accumulator func(traits.Mappable) bool) {
	store := context.KVStore(mapper.storeKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, mapper.keyGenerator(id))

	defer kvStorePrefixIterator.Close()
	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		mappable := mapper.mappablePrototype()
		Error := mapper.codec.UnmarshalBinaryBare(kvStorePrefixIterator.Value(), &mappable)
		if Error != nil {
			panic(Error)
		}
		if accumulator(mappable) {
			break
		}
	}
}
func (mapper mapper) InitializeMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) utilities.Mapper {
	mapper.storeKey = storeKey
	mapper.codec = codec
	return mapper
}
func (mapper mapper) RegisterCodec(codec *codec.Codec) {
	mapper.registerCodec(codec)
}

func NewMapper(keyGenerator func(types.ID) []byte, mappablePrototype func() traits.Mappable, registerCodec func(*codec.Codec)) *mapper {
	return &mapper{keyGenerator: keyGenerator, mappablePrototype: mappablePrototype, registerCodec: registerCodec}
}
