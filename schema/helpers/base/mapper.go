package base

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type mapper struct {
	storeKey          sdkTypes.StoreKey
	keyGenerator      func(types.ID) []byte
	mappablePrototype func() traits.Mappable
	registerCodec     func(*codec.Codec)
}

var _ helpers.Mapper = (*mapper)(nil)

func (mapper mapper) Create(context sdkTypes.Context, mappable traits.Mappable) {
	bytes := mappable.Encode()
	kvStore := context.KVStore(mapper.storeKey)
	kvStore.Set(mapper.keyGenerator(mappable.GetID()), bytes)
}
func (mapper mapper) Read(context sdkTypes.Context, id types.ID) traits.Mappable {
	kvStore := context.KVStore(mapper.storeKey)
	bytes := kvStore.Get(mapper.keyGenerator(id))
	if bytes == nil {
		return nil
	}
	return mapper.mappablePrototype().Decode(bytes)
}
func (mapper mapper) Update(context sdkTypes.Context, mappable traits.Mappable) {
	bytes := mappable.Encode()
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
		mappable := mapper.mappablePrototype().Decode(kvStorePrefixIterator.Value())
		if accumulator(mappable) {
			break
		}
	}
}
func (mapper mapper) RegisterCodec(codec *codec.Codec) {
	mapper.registerCodec(codec)
}

func NewMapper(module string, keyGenerator func(types.ID) []byte, mappablePrototype func() traits.Mappable, registerCodec func(*codec.Codec)) helpers.Mapper {
	return mapper{
		storeKey:          sdkTypes.NewKVStoreKey(module),
		keyGenerator:      keyGenerator,
		mappablePrototype: mappablePrototype,
		registerCodec:     registerCodec,
	}
}
