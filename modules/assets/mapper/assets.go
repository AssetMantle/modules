package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

type assets struct {
	ID   types.ID
	List []types.InterNFT

	mapper  assetsMapper
	context sdkTypes.Context
}

var _ types.InterNFTs = (*assets)(nil)

func (assets assets) GetID() types.ID { return assets.ID }
func (assets assets) Get(id types.ID) types.InterNFT {
	assetID := assetIDFromInterface(id)
	for _, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(assetID) == 0 {
			return oldAsset
		}
	}
	return nil
}
func (assets assets) GetList() []types.InterNFT {
	return assets.List
}

func (assets assets) Fetch(id types.ID) types.InterNFTs {
	var assetList []types.InterNFT
	assetsID := assetIDFromInterface(id)
	if len(assetsID.HashID.Bytes()) > 0 {
		asset := assets.mapper.read(assets.context, assetsID)
		if asset != nil {
			assetList = append(assetList, asset)
		}
	} else {
		appendAssetList := func(asset types.InterNFT) bool {
			assetList = append(assetList, asset)
			return false
		}
		assets.mapper.iterate(assets.context, assetsID, appendAssetList)
	}
	assets.ID, assets.List = id, assetList
	return assets
}
func (assets assets) Add(asset types.InterNFT) types.InterNFTs {
	assets.ID = readAssetID("")
	assets.mapper.create(assets.context, asset)
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) < 0 {
			assets.List = append(append(assets.List[:i], asset), assets.List[i+1:]...)
			break
		}
	}
	return assets
}
func (assets assets) Remove(asset types.InterNFT) types.InterNFTs {
	assets.mapper.delete(assets.context, asset.GetID())
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			assets.List = append(assets.List[:i], assets.List[i+1:]...)
			break
		}
	}
	return assets
}
func (assets assets) Mutate(asset types.InterNFT) types.InterNFTs {
	assets.mapper.update(assets.context, asset)
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			assets.List[i] = asset
			break
		}
	}
	return assets
}

func NewAssets(Mapper types.Mapper, context sdkTypes.Context) types.InterNFTs {
	switch mapper := Mapper.(type) {
	case assetsMapper:
		return assets{
			ID:      readAssetID(""),
			List:    []types.InterNFT{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleName)))
	}

}
