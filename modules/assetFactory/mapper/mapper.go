package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

func storeKey(assetID assetID) []byte {
	return append(constants.StoreKeyPrefix, assetID.Bytes()...)
}

type Mapper interface {
	create(sdkTypes.Context, baseAsset)
	read(sdkTypes.Context, assetID) baseAsset
	update(sdkTypes.Context, baseAsset)
	delete(sdkTypes.Context, assetID)
	iterate(sdkTypes.Context, assetID, func(baseAsset) bool)

	assetBaseImplementationFromInterface(asset types.Asset) baseAsset

	GenerateAssetID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) assetID

	New(sdkTypes.Context) types.Assets
	Assets(sdkTypes.Context, assetID) types.Assets
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
	kvStore.Set(storeKey(baseAsset.assetID), bytes)
}
func (baseMapper baseMapper) read(context sdkTypes.Context, assetID assetID) baseAsset {
	kvStore := context.KVStore(baseMapper.storeKey)
	bytes := kvStore.Get(storeKey(assetID))
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
	assetID := baseAsset.assetID
	kvStore := context.KVStore(baseMapper.storeKey)
	kvStore.Set(storeKey(assetID), bytes)
}
func (baseMapper baseMapper) delete(context sdkTypes.Context, assetID assetID) {
	bytes, err := baseMapper.codec.MarshalBinaryBare(&baseAsset{})
	if err != nil {
		panic(err)
	}
	kvStore := context.KVStore(baseMapper.storeKey)
	kvStore.Set(storeKey(assetID), bytes)
}
func (baseMapper baseMapper) iterate(context sdkTypes.Context, assetID assetID, accumulator func(baseAsset) bool) {
	store := context.KVStore(baseMapper.storeKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, storeKey(assetID))

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

func (baseMapper baseMapper) assetBaseImplementationFromInterface(asset types.Asset) baseAsset {
	return baseAsset{
		assetID{
			asset.ChainID(),
			asset.MaintainersID(),
			asset.ClassificationID(),
			asset.HashID(),
		},
		asset.Properties(),
		asset.GetLock(),
		asset.GetBurn(),
	}
}

func (baseMapper baseMapper) GenerateAssetID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) assetID {
	return assetID{chainID, maintainersID, classificationID, hashID}
}

func (baseMapper baseMapper) New(context sdkTypes.Context) types.Assets {
	return &baseAssets{baseMapper: baseMapper, context: context}
}

func (baseMapper baseMapper) Assets(context sdkTypes.Context, assetID assetID) types.Assets {
	var baseAssetList []baseAsset

	appendBaseAssetList := func(baseAsset baseAsset) bool {
		baseAssetList = append(baseAssetList, baseAsset)
		return false
	}
	baseMapper.iterate(context, assetID, appendBaseAssetList)

	return &baseAssets{assetID, baseAssetList, baseMapper, context}
}
