package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

func storeKey(identityID types.ID) []byte {
	return append(constants.StoreKeyPrefix, identityIDFromInterface(identityID).Bytes()...)
}

type identitiesMapper interface {
	types.Mapper
	create(sdkTypes.Context, types.InterIdentity)
	read(sdkTypes.Context, types.ID) types.InterIdentity
	update(sdkTypes.Context, types.InterIdentity)
	delete(sdkTypes.Context, types.ID)
	iterate(sdkTypes.Context, types.ID, func(types.InterIdentity) bool)
}

type mapper struct {
	StoreKey sdkTypes.StoreKey
	Codec    *codec.Codec
}

var _ identitiesMapper = (*mapper)(nil)

func (mapper mapper) InitializeMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) types.Mapper {
	mapper.StoreKey = storeKey
	mapper.Codec = codec
	return mapper
}
func (mapper mapper) create(context sdkTypes.Context, identity types.InterIdentity) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(identity)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(identity.GetID()), bytes)
}
func (mapper mapper) read(context sdkTypes.Context, identityID types.ID) types.InterIdentity {
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
func (mapper mapper) update(context sdkTypes.Context, identity types.InterIdentity) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(identity)
	if Error != nil {
		panic(Error)
	}
	identityID := identity.GetID()
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(identityID), bytes)
}
func (mapper mapper) delete(context sdkTypes.Context, identityID types.ID) {
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Delete(storeKey(identityID))
}
func (mapper mapper) iterate(context sdkTypes.Context, identityID types.ID, accumulator func(types.InterIdentity) bool) {
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

func newMapper() types.Mapper {
	return mapper{}
}

var Mapper = newMapper()
