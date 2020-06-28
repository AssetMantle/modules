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

func storeKey(id types.ID) []byte {
	return append(constants.StoreKeyPrefix, id.Bytes()...)
}

type Mapper interface {
	create(sdkTypes.Context, types.InterNFT)
	read(sdkTypes.Context, types.ID) types.InterNFT
	update(sdkTypes.Context, types.InterNFT)
	delete(sdkTypes.Context, types.ID)
	iterate(sdkTypes.Context, types.ID, func(types.InterNFT) bool)

	//Move to assetID
	MakeHashID(immutablePropertyList []types.Property) types.ID

	New(sdkTypes.Context) types.InterNFTs
	Assets(sdkTypes.Context, types.ID) types.InterNFTs
}

type mapper struct {
	storeKey sdkTypes.StoreKey
	codec    *codec.Codec
}

var _ Mapper = (*mapper)(nil)

func (mapper mapper) create(context sdkTypes.Context, asset types.InterNFT) {
	bytes, Error := mapper.codec.MarshalBinaryBare(asset)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.storeKey)
	kvStore.Set(storeKey(asset.GetID()), bytes)
}
func (mapper mapper) read(context sdkTypes.Context, assetID types.ID) types.InterNFT {
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
func (mapper mapper) update(context sdkTypes.Context, asset types.InterNFT) {
	bytes, Error := mapper.codec.MarshalBinaryBare(asset)
	if Error != nil {
		panic(Error)
	}
	assetID := asset.GetID()
	kvStore := context.KVStore(mapper.storeKey)
	kvStore.Set(storeKey(assetID), bytes)
}
func (mapper mapper) delete(context sdkTypes.Context, assetID types.ID) {
	bytes, Error := mapper.codec.MarshalBinaryBare(&asset{})
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.storeKey)
	kvStore.Set(storeKey(assetID), bytes)
}
func (mapper mapper) iterate(context sdkTypes.Context, assetID types.ID, accumulator func(types.InterNFT) bool) {
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
		facts = append(facts, immutableProperty.GetFact().String())
	}
	sort.Strings(facts)
	toDigest := strings.Join(facts, constants.PropertySeparator)
	h := sha1.New()
	h.Write([]byte(toDigest))
	return types.NewID(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}

func (mapper mapper) New(context sdkTypes.Context) types.InterNFTs {
	return &assets{Mapper: mapper, Context: context}
}

func (mapper mapper) Assets(context sdkTypes.Context, id types.ID) types.InterNFTs {
	var assetList []types.InterNFT

	appendAssetList := func(asset types.InterNFT) bool {
		assetList = append(assetList, asset)
		return false
	}
	assetID := assetIDFromInterface(id)
	mapper.iterate(context, assetID, appendAssetList)

	return &assets{assetID, assetList, mapper, context}
}

func NewMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) Mapper {
	return mapper{
		storeKey: storeKey,
		codec:    codec,
	}
}
