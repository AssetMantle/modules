package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.InterNFTs = (*assets)(nil)

type assets struct {
	ID   types.ID
	List []types.InterNFT

	Mapper  mapper
	Context sdkTypes.Context
}

func (assets assets) GetID() types.ID { return assets.ID }
func (assets assets) Get(id types.ID) types.InterNFT {
	for _, asset := range assets.List {
		if asset.GetID().Compare(id) == 0 {
			return asset
		}
	}
	return nil
}

func (assets assets) Add(asset types.InterNFT) types.InterNFTs {
	assets.Mapper.create(assets.Context, asset)
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) < 0 {
			assets.List = append(append(assets.List[:i], asset), assets.List[i+1:]...)
			break
		}
	}
	return assets
}
func (assets assets) Remove(asset types.InterNFT) types.InterNFTs {
	assets.Mapper.delete(assets.Context, asset.GetID())
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			assets.List = append(assets.List[:i], assets.List[i+1:]...)
			break
		}
	}
	return assets
}
func (assets assets) Mutate(asset types.InterNFT) types.InterNFTs {
	assets.Mapper.update(assets.Context, asset)
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			assets.List[i] = asset
			break
		}
	}
	return assets
}
