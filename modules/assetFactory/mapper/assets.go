package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.InterNFTs = (*assets)(nil)

type assets struct {
	assetID   assetID
	assetList []asset

	mapper  mapper
	context sdkTypes.Context
}

func (assets assets) ID() types.ID { return assets.assetID }
func (assets assets) Get(id types.ID) types.InterNFT {
	for _, asset := range assets.assetList {
		if asset.assetID.Compare(id) == 0 {
			return &asset
		}
	}
	return nil
}

func (assets *assets) Add(interNFT types.InterNFT) error {
	asset := assets.mapper.assetFromInterNFT(interNFT)
	assets.mapper.create(assets.context, asset)
	for i, oldAsset := range assets.assetList {
		if oldAsset.ID().Compare(asset.ID()) < 0 {
			assets.assetList = append(append(assets.assetList[:i], asset), assets.assetList[i+1:]...)
			break
		}
	}
	return nil
}
func (assets *assets) Remove(interNFT types.InterNFT) error {
	assetID := assets.mapper.assetFromInterNFT(interNFT).assetID
	assets.mapper.delete(assets.context, assetID)
	for i, asset := range assets.assetList {
		if asset.ID().Compare(assetID) == 0 {
			assets.assetList = append(assets.assetList[:i], assets.assetList[i+1:]...)
			break
		}
	}
	return nil
}
func (assets *assets) Mutate(interNFT types.InterNFT) error {
	asset := assets.mapper.assetFromInterNFT(interNFT)
	assets.mapper.update(assets.context, asset)
	for i, oldAsset := range assets.assetList {
		if oldAsset.ID().Compare(asset.ID()) == 0 {
			assets.assetList[i] = asset
			break
		}
	}
	return nil
}
