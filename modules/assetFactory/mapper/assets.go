package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.InterNFTs = (*assets)(nil)

type assets struct {
	AssetID   assetID
	AssetList []asset

	Mapper  mapper
	Context sdkTypes.Context
}

func (assets assets) ID() types.ID { return assets.AssetID }
func (assets assets) Get(id types.ID) types.InterNFT {
	for _, asset := range assets.AssetList {
		if asset.AssetID.Compare(id) == 0 {
			return &asset
		}
	}
	return nil
}

func (assets *assets) Add(interNFT types.InterNFT) error {
	asset := assetFromInterface(interNFT)
	assets.Mapper.create(assets.Context, asset)
	for i, oldAsset := range assets.AssetList {
		if oldAsset.ID().Compare(asset.ID()) < 0 {
			assets.AssetList = append(append(assets.AssetList[:i], asset), assets.AssetList[i+1:]...)
			break
		}
	}
	return nil
}
func (assets *assets) Remove(interNFT types.InterNFT) error {
	assetID := assetFromInterface(interNFT).AssetID
	assets.Mapper.delete(assets.Context, assetID)
	for i, asset := range assets.AssetList {
		if asset.ID().Compare(assetID) == 0 {
			assets.AssetList = append(assets.AssetList[:i], assets.AssetList[i+1:]...)
			break
		}
	}
	return nil
}
func (assets *assets) Mutate(interNFT types.InterNFT) error {
	asset := assetFromInterface(interNFT)
	assets.Mapper.update(assets.Context, asset)
	for i, oldAsset := range assets.AssetList {
		if oldAsset.ID().Compare(asset.ID()) == 0 {
			assets.AssetList[i] = asset
			break
		}
	}
	return nil
}
