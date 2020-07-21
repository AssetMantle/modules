package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

func storeKey(identityID schema.ID) []byte {
	return append(StoreKeyPrefix, identityIDFromInterface(identityID).Bytes()...)
}

type identitiesMapper interface {
	utility.Mapper
	create(sdkTypes.Context, schema.InterIdentity)
	read(sdkTypes.Context, schema.ID) schema.InterIdentity
	update(sdkTypes.Context, schema.InterIdentity)
	delete(sdkTypes.Context, schema.ID)
	iterate(sdkTypes.Context, schema.ID, func(schema.InterIdentity) bool)
}

type mapper struct {
	StoreKey sdkTypes.StoreKey
	Codec    *codec.Codec
}

var _ identitiesMapper = (*mapper)(nil)

func (mapper mapper) InitializeMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) utility.Mapper {
	mapper.StoreKey = storeKey
	mapper.Codec = codec
	return mapper
}
func (mapper mapper) create(context sdkTypes.Context, identity schema.InterIdentity) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(identity)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(identity.GetID()), bytes)
}
func (mapper mapper) read(context sdkTypes.Context, identityID schema.ID) schema.InterIdentity {
	kvStore := context.KVStore(mapper.StoreKey)
	bytes := kvStore.Get(storeKey(identityID))
	if bytes == nil {
		return nil
	}
	identity := identity{}
	Error := mapper.Codec.UnmarshalBinaryBare(bytes, &identity)
	if Error != nil {
		panic(Error)
	}
	return identity
}
func (mapper mapper) update(context sdkTypes.Context, identity schema.InterIdentity) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(identity)
	if Error != nil {
		panic(Error)
	}
	identityID := identity.GetID()
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(identityID), bytes)
}
func (mapper mapper) delete(context sdkTypes.Context, identityID schema.ID) {
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Delete(storeKey(identityID))
}
func (mapper mapper) iterate(context sdkTypes.Context, identityID schema.ID, accumulator func(schema.InterIdentity) bool) {
	store := context.KVStore(mapper.StoreKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, storeKey(identityID))

	defer kvStorePrefixIterator.Close()
	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		identity := identity{}
		Error := mapper.Codec.UnmarshalBinaryBare(kvStorePrefixIterator.Value(), &identity)
		if Error != nil {
			panic(Error)
		}
		if accumulator(identity) {
			break
		}
	}
}

func newMapper() utility.Mapper {
	return mapper{}
}

var Mapper = newMapper()
