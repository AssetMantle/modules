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
	create(sdkTypes.Context, asset)
	read(sdkTypes.Context, assetID) asset
	update(sdkTypes.Context, asset)
	delete(sdkTypes.Context, assetID)
	iterate(sdkTypes.Context, assetID, func(asset) bool)

	MakeHashID(immutablePropertyList []types.Property) types.ID
	MakeAssetID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) types.ID
	MakeAsset(assetID types.ID, properties types.Properties, lock types.Height, burn types.Height) types.InterNFT

	New(sdkTypes.Context) types.InterNFTs
	Assets(sdkTypes.Context, types.ID) types.InterNFTs
	QueryAssets(sdkTypes.Context, types.ID) []byte
}

type mapper struct {
	storeKey sdkTypes.StoreKey
	codec    *codec.Codec
}

func NewMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) Mapper {
	return mapper{
		storeKey: storeKey,
		codec:    codec,
	}
}

var _ Mapper = (*mapper)(nil)

func (mapper mapper) create(context sdkTypes.Context, asset asset) {
	bytes, Error := mapper.codec.MarshalBinaryBare(asset)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.storeKey)
	kvStore.Set(storeKey(asset.AssetID), bytes)
}
func (mapper mapper) read(context sdkTypes.Context, assetID assetID) asset {
	kvStore := context.KVStore(mapper.storeKey)
	bytes := kvStore.Get(storeKey(assetID))
	if bytes == nil {
		return asset{}
	}
	asset := asset{}
	Error := mapper.codec.UnmarshalBinaryBare(bytes, &asset)
	if Error != nil {
		panic(Error)
	}
	return asset
}
func (mapper mapper) update(context sdkTypes.Context, asset asset) {
	bytes, Error := mapper.codec.MarshalBinaryBare(asset)
	if Error != nil {
		panic(Error)
	}
	assetID := asset.AssetID
	kvStore := context.KVStore(mapper.storeKey)
	kvStore.Set(storeKey(assetID), bytes)
}
func (mapper mapper) delete(context sdkTypes.Context, assetID assetID) {
	bytes, Error := mapper.codec.MarshalBinaryBare(&asset{})
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.storeKey)
	kvStore.Set(storeKey(assetID), bytes)
}
func (mapper mapper) iterate(context sdkTypes.Context, assetID assetID, accumulator func(asset) bool) {
	store := context.KVStore(mapper.storeKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, storeKey(assetID))

	defer kvStorePrefixIterator.Close()
	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		asset := asset{}
		Error := mapper.codec.UnmarshalBinaryBare(kvStorePrefixIterator.Value(), &asset)
		if Error != nil {
			panic(Error)
		}
		if accumulator(asset) {
			break
		}
	}
}

func (mapper mapper) MakeHashID(immutablePropertyList []types.Property) types.ID {
	var facts []string
	for _, immutableProperty := range immutablePropertyList {
		facts = append(facts, immutableProperty.Fact().String())
	}
	sort.Strings(facts)
	toDigest := strings.Join(facts, constants.PropertySeparator)
	h := sha1.New()
	h.Write([]byte(toDigest))
	return types.BaseID{IDString: base64.URLEncoding.EncodeToString(h.Sum(nil))}
}

func (mapper mapper) MakeAssetID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) types.ID {
	return assetID{chainID, maintainersID, classificationID, hashID}
}

func (mapper mapper) MakeAsset(id types.ID, properties types.Properties, lock types.Height, burn types.Height) types.InterNFT {
	assetID := assetIDFromInterface(id)
	return &asset{AssetID: assetID, BaseProperties: types.BasePropertiesFromInterface(properties), Lock: types.BaseHeightFromInterface(lock), Burn: types.BaseHeightFromInterface(burn)}
}

func (mapper mapper) New(context sdkTypes.Context) types.InterNFTs {
	return &assets{Mapper: mapper, Context: context}
}

func (mapper mapper) Assets(context sdkTypes.Context, id types.ID) types.InterNFTs {
	var assetList []asset

	appendAssetList := func(asset asset) bool {
		assetList = append(assetList, asset)
		return false
	}
	assetID := assetIDFromInterface(id)
	mapper.iterate(context, assetID, appendAssetList)

	return &assets{assetID, assetList, mapper, context}
}
func (mapper mapper) QueryAssets(context sdkTypes.Context, id types.ID) []byte {
	bz, err := codec.MarshalJSONIndent(mapper.codec, mapper.Assets(context, id))
	if err != nil {
		return nil
	}
	return bz
}
