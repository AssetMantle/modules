package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type assets struct {
	ID   types.ID
	List []entities.InterNFT

	mapper  assetsMapper
	context sdkTypes.Context
}

var _ mappers.InterNFTs = (*assets)(nil)

func (assets assets) GetID() types.ID { return assets.ID }
func (assets assets) Get(id types.ID) entities.InterNFT {
	assetID := assetIDFromInterface(id)
	for _, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(assetID) == 0 {
			return oldAsset
		}
	}
	return nil
}
func (assets assets) GetList() []entities.InterNFT {
	return assets.List
}

func (assets assets) Fetch(id types.ID) mappers.InterNFTs {
	var assetList []entities.InterNFT
	assetsID := assetIDFromInterface(id)
	if len(assetsID.HashID.Bytes()) > 0 {
		asset := assets.mapper.read(assets.context, assetsID)
		if asset != nil {
			assetList = append(assetList, asset)
		}
	} else {
		appendAssetList := func(asset entities.InterNFT) bool {
			assetList = append(assetList, asset)
			return false
		}
		assets.mapper.iterate(assets.context, assetsID, appendAssetList)
	}
	assets.ID, assets.List = id, assetList
	return assets
}
func (assets assets) Add(asset entities.InterNFT) mappers.InterNFTs {
	assets.ID = readAssetID("")
	assets.mapper.create(assets.context, asset)
	assets.List = append(assets.List, asset)
	return assets
}
func (assets assets) Remove(asset entities.InterNFT) mappers.InterNFTs {
	assets.mapper.delete(assets.context, asset.GetID())
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			assets.List = append(assets.List[:i], assets.List[i+1:]...)
			break
		}
	}
	return assets
}
func (assets assets) Mutate(asset entities.InterNFT) mappers.InterNFTs {
	assets.mapper.update(assets.context, asset)
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			assets.List[i] = asset
			break
		}
	}
	return assets
}

func NewAssets(Mapper utilities.Mapper, context sdkTypes.Context) mappers.InterNFTs {
	switch mapper := Mapper.(type) {
	case assetsMapper:
		return assets{
			ID:      readAssetID(""),
			List:    []entities.InterNFT{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleName)))
	}

}
