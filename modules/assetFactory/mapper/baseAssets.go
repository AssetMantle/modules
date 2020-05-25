package mapper

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.Assets = (*baseAssets)(nil)

type baseAssets struct {
	id   baseAssetID
	list []baseAsset
}

func (baseAssets baseAssets) String() string {
	bytes, Error := json.Marshal(baseAssets)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (baseAssets baseAssets) ID() types.ID { return baseAssets.id }
func (baseAssets baseAssets) Asset(id types.ID) types.Asset {
	for _, asset := range baseAssets.list {
		if asset.id.IsEqualTo(id) {
			return &asset
		}
	}
}

func (baseAssets *baseAssets) Add(types.Asset) error    {}
func (baseAssets *baseAssets) Remove(types.Asset) error {}
func (baseAssets *baseAssets) Mutate(types.Asset) error {}
