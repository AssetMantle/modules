package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

func storeKey(baseAssetID baseAssetID) []byte {
	return append(constants.StoreKeyPrefix, baseAssetID.Bytes()...)
}

type Mapper interface {
	create(sdkTypes.Context, baseAsset)
	read(sdkTypes.Context, baseAssetID) baseAsset
	update(sdkTypes.Context, baseAsset)
	delete(sdkTypes.Context, baseAssetID)
	iterate(sdkTypes.Context, baseAssetID, func(baseAsset) bool)

	assetBaseImplementationFromInterface(asset types.Asset) baseAsset

	AssetID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) baseAssetID

	New(sdkTypes.Context) types.Assets
	Assets(sdkTypes.Context, baseAssetID) types.Assets
}

type baseMapper struct {
	storeKey sdkTypes.StoreKey
	codec    *codec.Codec
}

func NewMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) Mapper {
	return baseMapper{
		storeKey: storeKey,
		codec:    codec,
	}
}

var _ Mapper = (*baseMapper)(nil)

func (baseMapper baseMapper) create(context sdkTypes.Context, baseAsset baseAsset) {
	bytes, err := baseMapper.codec.MarshalBinaryBare(baseAsset)
	if err != nil {
		panic(err)
	}
	kvStore := context.KVStore(baseMapper.storeKey)
	kvStore.Set(storeKey(baseAsset.baseAssetID), bytes)
}
func (baseMapper baseMapper) read(context sdkTypes.Context, baseAssetID baseAssetID) baseAsset {
	kvStore := context.KVStore(baseMapper.storeKey)
	bytes := kvStore.Get(storeKey(baseAssetID))
	if bytes == nil {
		return baseAsset{}
	}
	baseAsset := baseAsset{}
	err := baseMapper.codec.UnmarshalBinaryBare(bytes, &baseAsset)
	if err != nil {
		panic(err)
	}
	return baseAsset
}
func (baseMapper baseMapper) update(context sdkTypes.Context, baseAsset baseAsset) {
	bytes, err := baseMapper.codec.MarshalBinaryBare(baseAsset)
	if err != nil {
		panic(err)
	}
	baseAssetID := baseAsset.baseAssetID
	kvStore := context.KVStore(baseMapper.storeKey)
	kvStore.Set(storeKey(baseAssetID), bytes)
}
func (baseMapper baseMapper) delete(context sdkTypes.Context, baseAssetID baseAssetID) {
	bytes, err := baseMapper.codec.MarshalBinaryBare(&baseAsset{})
	if err != nil {
		panic(err)
	}
	kvStore := context.KVStore(baseMapper.storeKey)
	kvStore.Set(storeKey(baseAssetID), bytes)
}
func (baseMapper baseMapper) iterate(context sdkTypes.Context, baseAssetID baseAssetID, accumulator func(baseAsset) bool) {
	store := context.KVStore(baseMapper.storeKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, storeKey(baseAssetID))

	defer kvStorePrefixIterator.Close()
	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		baseAsset := baseAsset{}
		Error := baseMapper.codec.UnmarshalBinaryBare(kvStorePrefixIterator.Value(), &baseAsset)
		if Error != nil {
			panic(Error)
		}
		if accumulator(baseAsset) {
			break
		}
	}
}

func (baseMapper baseMapper) New(context sdkTypes.Context) types.Assets {
	return &baseAssets{baseMapper: baseMapper, context: context}
}

func (baseMapper baseMapper) Assets(context sdkTypes.Context, baseAssetID baseAssetID) types.Assets {
	var baseAssetList []baseAsset

	appendBaseAssetList := func(baseAsset baseAsset) bool {
		baseAssetList = append(baseAssetList, baseAsset)
		return false
	}
	baseMapper.iterate(context, baseAssetID, appendBaseAssetList)

	return &baseAssets{baseAssetID, baseAssetList, baseMapper, context}
}
func (baseMapper baseMapper) assetBaseImplementationFromInterface(asset types.Asset) baseAsset {
	return baseAsset{
		baseAssetID{
			asset.ChainID(),
			asset.MaintainersID(),
			asset.ClassificationID(),
			asset.HashID(),
		},
		asset.OwnersID(),
		asset.Properties(),
		asset.GetLock(),
		asset.GetBurn(),
	}
}

func (baseMapper baseMapper) AssetID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) baseAssetID {
	return baseAssetID{chainID, maintainersID, classificationID, hashID}
}
