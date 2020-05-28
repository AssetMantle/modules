package mapper

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"sort"
	"strings"
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

	GenerateHashID(immutablePropertyList []types.Property) types.ID
	GenerateAssetID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) types.ID
	MakeAsset(assetID types.ID, properties types.Properties, lock types.Height, burn types.Height) types.Asset

	New(sdkTypes.Context) types.Assets
	Assets(sdkTypes.Context, types.ID) types.Assets
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

func (baseMapper baseMapper) assetIDFromInterface(id types.ID) assetID {
	base64IDs := strings.Split(id.String(), constants.IDSeparator)
	chainID, _ := base64.URLEncoding.DecodeString(base64IDs[0])
	classificationID, _ := base64.URLEncoding.DecodeString(base64IDs[1])
	maintainersID, _ := base64.URLEncoding.DecodeString(base64IDs[2])
	hashID, _ := base64.URLEncoding.DecodeString(base64IDs[3])

	return assetID{
		chainID:          types.BaseID{BaseBytes: chainID},
		maintainersID:    types.BaseID{BaseBytes: classificationID},
		classificationID: types.BaseID{BaseBytes: maintainersID},
		hashID:           types.BaseID{BaseBytes: hashID},
	}
}

func (baseMapper baseMapper) GenerateHashID(immutablePropertyList []types.Property) types.ID {
	var facts []string
	for _, immutableProperty := range immutablePropertyList {
		facts = append(facts, immutableProperty.String())
	}
	sort.Strings(facts)
	toDigest := strings.Join(facts, constants.PropertySeparator)
	h := sha1.New()
	h.Write([]byte(toDigest))
	return types.BaseID{BaseBytes: h.Sum(nil)}
}

func (baseMapper baseMapper) GenerateAssetID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) types.ID {
	return assetID{chainID, maintainersID, classificationID, hashID}
}

func (baseMapper baseMapper) MakeAsset(id types.ID, properties types.Properties, lock types.Height, burn types.Height) types.Asset {
	assetID := baseMapper.assetIDFromInterface(id)
	return &baseAsset{assetID: assetID, properties: properties, lock: lock, burn: burn}
}

func (baseMapper baseMapper) New(context sdkTypes.Context) types.Assets {
	return &baseAssets{baseMapper: baseMapper, context: context}
}

func (baseMapper baseMapper) Assets(context sdkTypes.Context, id types.ID) types.Assets {
	var baseAssetList []baseAsset

	appendBaseAssetList := func(baseAsset baseAsset) bool {
		baseAssetList = append(baseAssetList, baseAsset)
		return false
	}
	assetID := baseMapper.assetIDFromInterface(id)
	baseMapper.iterate(context, assetID, appendBaseAssetList)

	return &baseAssets{assetID, baseAssetList, baseMapper, context}
}
