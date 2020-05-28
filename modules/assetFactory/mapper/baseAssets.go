package mapper

import (
	"encoding/json"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.Assets = (*baseAssets)(nil)

type baseAssets struct {
	assetID       assetID
	baseAssetList []baseAsset

	baseMapper baseMapper
	context    sdkTypes.Context
}

func (baseAssets baseAssets) String() string {
	bytes, Error := json.Marshal(baseAssets)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (baseAssets baseAssets) ID() types.ID { return baseAssets.assetID }
func (baseAssets baseAssets) Asset(id types.ID) types.Asset {
	for _, baseAsset := range baseAssets.baseAssetList {
		if baseAsset.assetID.Compare(id) == 0 {
			return &baseAsset
		}
	}
	return nil
}

func (baseAssets *baseAssets) Add(asset types.Asset) error {
	for i, baseAsset := range baseAssets.baseAssetList {
		if baseAsset.ID().Compare(asset.ID()) < 0 {
			baseAssets.baseAssetList = append(append(baseAssets.baseAssetList[:i], baseAsset), baseAssets.baseAssetList[i+1:]...)
			baseAssets.baseMapper.create(baseAssets.context, baseAssets.baseMapper.assetBaseImplementationFromInterface(asset))
			break
		}
	}
	return nil
}
func (baseAssets *baseAssets) Remove(asset types.Asset) error {
	for i, baseAsset := range baseAssets.baseAssetList {
		if baseAsset.ID().Compare(asset.ID()) == 0 {
			baseAssets.baseAssetList = append(baseAssets.baseAssetList[:i], baseAssets.baseAssetList[i+1:]...)
			baseAssets.baseMapper.delete(baseAssets.context, baseAssets.baseMapper.assetBaseImplementationFromInterface(asset).assetID)
			break
		}
	}
	return nil
}
func (baseAssets *baseAssets) Mutate(asset types.Asset) error {
	for i, baseAsset := range baseAssets.baseAssetList {
		if baseAsset.ID().Compare(asset.ID()) == 0 {
			baseAssets.baseAssetList[i] = baseAssets.baseMapper.assetBaseImplementationFromInterface(asset)
			baseAssets.baseMapper.update(baseAssets.context, baseAssets.baseMapper.assetBaseImplementationFromInterface(asset))
			break
		}
	}
	return nil
}
