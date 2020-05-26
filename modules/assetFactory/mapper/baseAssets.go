package mapper

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.Assets = (*baseAssets)(nil)

type baseAssets struct {
	baseAssetID   baseAssetID
	baseAssetList []baseAsset
}

func (baseAssets baseAssets) String() string {
	bytes, Error := json.Marshal(baseAssets)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (baseAssets baseAssets) ID() types.ID { return baseAssets.baseAssetID }
func (baseAssets baseAssets) Asset(id types.ID) types.Asset {
	for _, baseAsset := range baseAssets.baseAssetList {
		if baseAsset.baseAssetID.Compare(id) == 0 {
			return &baseAsset
		}
	}
	return nil
}

func (baseAssets *baseAssets) Add(asset types.Asset) error    {}
func (baseAssets *baseAssets) Remove(asset types.Asset) error {}
func (baseAssets *baseAssets) Mutate(asset types.Asset) error {}
