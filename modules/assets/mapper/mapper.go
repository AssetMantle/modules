package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

func storeKey(assetID schema.ID) []byte {
	return append(StoreKeyPrefix, assetIDFromInterface(assetID).Bytes()...)
}

type assetsMapper interface {
	utility.Mapper
	create(sdkTypes.Context, schema.InterNFT)
	read(sdkTypes.Context, schema.ID) schema.InterNFT
	update(sdkTypes.Context, schema.InterNFT)
	delete(sdkTypes.Context, schema.ID)
	iterate(sdkTypes.Context, schema.ID, func(schema.InterNFT) bool)
}

type mapper struct {
	StoreKey sdkTypes.StoreKey
	Codec    *codec.Codec
}

var _ assetsMapper = (*mapper)(nil)

func (mapper mapper) InitializeMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) utility.Mapper {
	mapper.StoreKey = storeKey
	mapper.Codec = codec
	return mapper
}
func (mapper mapper) create(context sdkTypes.Context, asset schema.InterNFT) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(asset)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(asset.GetID()), bytes)
}
func (mapper mapper) read(context sdkTypes.Context, assetID schema.ID) schema.InterNFT {
	kvStore := context.KVStore(mapper.StoreKey)
	bytes := kvStore.Get(storeKey(assetID))
	if bytes == nil {
		return nil
	}
	asset := asset{}
	Error := mapper.Codec.UnmarshalBinaryBare(bytes, &asset)
	if Error != nil {
		panic(Error)
	}
	return asset
}
func (mapper mapper) update(context sdkTypes.Context, asset schema.InterNFT) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(asset)
	if Error != nil {
		panic(Error)
	}
	assetID := asset.GetID()
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(assetID), bytes)
}
func (mapper mapper) delete(context sdkTypes.Context, assetID schema.ID) {
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Delete(storeKey(assetID))
}
func (mapper mapper) iterate(context sdkTypes.Context, assetID schema.ID, accumulator func(schema.InterNFT) bool) {
	store := context.KVStore(mapper.StoreKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, storeKey(assetID))

	defer kvStorePrefixIterator.Close()
	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		asset := asset{}
		Error := mapper.Codec.UnmarshalBinaryBare(kvStorePrefixIterator.Value(), &asset)
		if Error != nil {
			panic(Error)
		}
		if accumulator(asset) {
			break
		}
	}
}

func newMapper() utility.Mapper {
	return mapper{}
}

var Mapper = newMapper()
