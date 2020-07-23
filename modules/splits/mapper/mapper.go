package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

func storeKey(splitID types.ID) []byte {
	return append(StoreKeyPrefix, splitIDFromInterface(splitID).Bytes()...)
}

//TODO make DAO interface
type splitsMapper interface {
	utilities.Mapper
	create(sdkTypes.Context, entities.Split)
	read(sdkTypes.Context, types.ID) entities.Split
	update(sdkTypes.Context, entities.Split)
	delete(sdkTypes.Context, types.ID)
	iterate(sdkTypes.Context, types.ID, func(entities.Split) bool)
}

type mapper struct {
	StoreKey sdkTypes.StoreKey
	Codec    *codec.Codec
}

var _ splitsMapper = (*mapper)(nil)

func (mapper mapper) InitializeMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) utilities.Mapper {
	mapper.StoreKey = storeKey
	mapper.Codec = codec
	return mapper
}
func (mapper mapper) create(context sdkTypes.Context, split entities.Split) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(split)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(split.GetID()), bytes)
}
func (mapper mapper) read(context sdkTypes.Context, splitID types.ID) entities.Split {
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
func (mapper mapper) update(context sdkTypes.Context, split entities.Split) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(split)
	if Error != nil {
		panic(Error)
	}
	splitID := split.GetID()
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(splitID), bytes)
}
func (mapper mapper) delete(context sdkTypes.Context, splitID types.ID) {
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Delete(storeKey(splitID))
}
func (mapper mapper) iterate(context sdkTypes.Context, splitID types.ID, accumulator func(entities.Split) bool) {
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

func newMapper() utilities.Mapper {
	return mapper{}
}

var Mapper = newMapper()
