package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type assets struct {
	ID   schema.ID
	List []schema.InterNFT

	mapper  assetsMapper
	context sdkTypes.Context
}

var _ schema.InterNFTs = (*assets)(nil)

func (assets assets) GetID() schema.ID { return assets.ID }
func (assets assets) Get(id schema.ID) schema.InterNFT {
	assetID := assetIDFromInterface(id)
	for _, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(assetID) == 0 {
			return oldAsset
		}
	}
	return nil
}
func (assets assets) GetList() []schema.InterNFT {
	return assets.List
}

func (assets assets) Fetch(id schema.ID) schema.InterNFTs {
	var assetList []schema.InterNFT
	assetsID := assetIDFromInterface(id)
	if len(assetsID.HashID.Bytes()) > 0 {
		asset := assets.mapper.read(assets.context, assetsID)
		if asset != nil {
			assetList = append(assetList, asset)
		}
	} else {
		appendAssetList := func(asset schema.InterNFT) bool {
			assetList = append(assetList, asset)
			return false
		}
		assets.mapper.iterate(assets.context, assetsID, appendAssetList)
	}
	assets.ID, assets.List = id, assetList
	return assets
}
func (assets assets) Add(asset schema.InterNFT) schema.InterNFTs {
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
func (assets assets) Remove(asset schema.InterNFT) schema.InterNFTs {
	assets.mapper.delete(assets.context, asset.GetID())
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			assets.List = append(assets.List[:i], assets.List[i+1:]...)
			break
		}
	}
	return assets
}
func (assets assets) Mutate(asset schema.InterNFT) schema.InterNFTs {
	assets.mapper.update(assets.context, asset)
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			assets.List[i] = asset
			break
		}
	}
	return assets
}

func NewAssets(Mapper utility.Mapper, context sdkTypes.Context) schema.InterNFTs {
	switch mapper := Mapper.(type) {
	case assetsMapper:
		return assets{
			ID:      readAssetID(""),
			List:    []schema.InterNFT{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleName)))
	}

}
