package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

func storeKey(assetID types.ID) []byte {
	return append(constants.StoreKeyPrefix, []byte("uncommentAftermakingThatFunction")...) //assetIDFromInterface(assetID).Bytes()...)
}

type assetsMapper interface {
	types.Mapper
	create(sdkTypes.Context, types.InterNFT)
	//read(sdkTypes.Context, types.ID) types.InterNFT
	update(sdkTypes.Context, types.InterNFT)
	delete(sdkTypes.Context, types.ID)
	//iterate(sdkTypes.Context, types.ID, func(types.InterNFT) bool)
}

type mapper struct {
	StoreKey sdkTypes.StoreKey
	Codec    *codec.Codec
}

var _ assetsMapper = (*mapper)(nil)

func (mapper mapper) InitializeMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) types.Mapper {
	mapper.StoreKey = storeKey
	mapper.Codec = codec
	return mapper
}
func (mapper mapper) create(context sdkTypes.Context, asset types.InterNFT) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(asset)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(asset.GetID()), bytes)
}

//func (mapper mapper) read(context sdkTypes.Context, assetID types.ID) types.InterNFT {
//	kvStore := context.KVStore(mapper.StoreKey)
//	bytes := kvStore.Get(storeKey(assetID))
//	if bytes == nil {
//		return nil
//	}
//	asset := asset{}
//	Error := mapper.Codec.UnmarshalBinaryBare(bytes, &asset)
//	if Error != nil {
//		panic(Error)
//	}
//	return asset
//}
func (mapper mapper) update(context sdkTypes.Context, asset types.InterNFT) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(asset)
	if Error != nil {
		panic(Error)
	}
	assetID := asset.GetID()
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(assetID), bytes)
}
func (mapper mapper) delete(context sdkTypes.Context, assetID types.ID) {
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Delete(storeKey(assetID))
}

//func (mapper mapper) iterate(context sdkTypes.Context, assetID types.ID, accumulator func(types.InterNFT) bool) {
//	store := context.KVStore(mapper.StoreKey)
//	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, storeKey(assetID))
//
//	defer kvStorePrefixIterator.Close()
//	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
//		asset := asset {}
//		Error := mapper.Codec.UnmarshalBinaryBare(kvStorePrefixIterator.Value(), &asset)
//		if Error != nil {
//			panic(Error)
//		}
//		if accumulator(asset) {
//			break
//		}
//	}
//}

func newMapper() types.Mapper {
	return mapper{}
}

var Mapper = newMapper()
