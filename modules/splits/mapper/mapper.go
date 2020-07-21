package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

func storeKey(splitID schema.ID) []byte {
	return append(StoreKeyPrefix, splitIDFromInterface(splitID).Bytes()...)
}

//TODO make DAO interface
type splitsMapper interface {
	utility.Mapper
	create(sdkTypes.Context, schema.Split)
	read(sdkTypes.Context, schema.ID) schema.Split
	update(sdkTypes.Context, schema.Split)
	delete(sdkTypes.Context, schema.ID)
	iterate(sdkTypes.Context, schema.ID, func(schema.Split) bool)
}

type mapper struct {
	StoreKey sdkTypes.StoreKey
	Codec    *codec.Codec
}

var _ splitsMapper = (*mapper)(nil)

func (mapper mapper) InitializeMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) utility.Mapper {
	mapper.StoreKey = storeKey
	mapper.Codec = codec
	return mapper
}
func (mapper mapper) create(context sdkTypes.Context, split schema.Split) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(split)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(split.GetID()), bytes)
}
func (mapper mapper) read(context sdkTypes.Context, splitID schema.ID) schema.Split {
	kvStore := context.KVStore(mapper.StoreKey)
	bytes := kvStore.Get(storeKey(splitID))
	if bytes == nil {
		return nil
	}
	split := split{}
	Error := mapper.Codec.UnmarshalBinaryBare(bytes, &split)
	if Error != nil {
		panic(Error)
	}
	return split
}
func (mapper mapper) update(context sdkTypes.Context, split schema.Split) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(split)
	if Error != nil {
		panic(Error)
	}
	splitID := split.GetID()
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(splitID), bytes)
}
func (mapper mapper) delete(context sdkTypes.Context, splitID schema.ID) {
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Delete(storeKey(splitID))
}
func (mapper mapper) iterate(context sdkTypes.Context, splitID schema.ID, accumulator func(schema.Split) bool) {
	store := context.KVStore(mapper.StoreKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, storeKey(splitID))

	defer kvStorePrefixIterator.Close()
	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		split := split{}
		Error := mapper.Codec.UnmarshalBinaryBare(kvStorePrefixIterator.Value(), &split)
		if Error != nil {
			panic(Error)
		}
		if accumulator(split) {
			break
		}
	}
}

func newMapper() utility.Mapper {
	return mapper{}
}

var Mapper = newMapper()
