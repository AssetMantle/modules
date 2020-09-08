/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type assets struct {
	ID   types.ID             `json:"id" valid:"required~required field id missing"`
	List []mappables.InterNFT `json:"list" valid:"required~required field list missing"`

	mapper  helpers.Mapper
	context sdkTypes.Context
}

var _ mappers.InterNFTs = (*assets)(nil)

func (assets assets) GetID() types.ID { return assets.ID }
func (assets assets) Get(id types.ID) mappables.InterNFT {
	assetID := assetIDFromInterface(id)
	for _, oldAsset := range assets.List {
		if oldAsset.GetID().Equal(assetID) {
			return oldAsset
		}
	}
	return nil
}
func (assets assets) GetList() []mappables.InterNFT {
	return assets.List
}

func (assets assets) Fetch(id types.ID) mappers.InterNFTs {
	var assetList []mappables.InterNFT
	assetsID := assetIDFromInterface(id)
	if len(assetsID.HashID.Bytes()) > 0 {
		mappable := assets.mapper.Read(assets.context, assetsID)
		if mappable != nil {
			assetList = append(assetList, mappable.(asset))
		}
	} else {
		appendMappableList := func(mappable traits.Mappable) bool {
			assetList = append(assetList, mappable.(asset))
			return false
		}
		assets.mapper.Iterate(assets.context, assetsID, appendMappableList)
	}
	assets.ID, assets.List = id, assetList
	return assets
}
func (assets assets) Add(asset mappables.InterNFT) mappers.InterNFTs {
	assets.ID = readAssetID("")
	assets.mapper.Create(assets.context, asset)
	assets.List = append(assets.List, asset)
	return assets
}
func (assets assets) Remove(asset mappables.InterNFT) mappers.InterNFTs {
	assets.mapper.Delete(assets.context, asset.GetID())
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Equal(asset.GetID()) {
			assets.List = append(assets.List[:i], assets.List[i+1:]...)
			break
		}
	}
	return assets
}
func (assets assets) Mutate(asset mappables.InterNFT) mappers.InterNFTs {
	assets.mapper.Update(assets.context, asset)
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Equal(asset.GetID()) {
			assets.List[i] = asset
			break
		}
	}
	return assets
}

func NewAssets(context sdkTypes.Context, mapper helpers.Mapper) mappers.InterNFTs {
	return assets{
		ID:      readAssetID(""),
		List:    []mappables.InterNFT{},
		mapper:  mapper,
		context: context,
	}
}
