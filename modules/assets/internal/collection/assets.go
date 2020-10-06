/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package collection

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type assets struct {
	Key  helpers.Key       `json:"key" valid:"required~required field key missing"`
	List []traits.Mappable `json:"list" valid:"required~required field list missing"`

	mapper  helpers.Mapper
	context sdkTypes.Context
}

var _ mappers.InterNFTs = (*assets)(nil)

func (assets assets) GetKey() helpers.Key { return assets.Key }
func (assets assets) Get(assetID types.ID) traits.Mappable {
	for _, oldAsset := range assets.List {
		if oldAsset.GetID().Equal(assetID) {
			return oldAsset
		}
	}
	return nil
}
func (assets assets) GetList() []traits.Mappable {
	var mappableList []traits.Mappable
	for _, asset := range assets.List {
		mappableList = append(mappableList, asset)
	}
	return mappableList
}

func (assets assets) Fetch(assetID types.ID) traits.Collection {
	var assetList []traits.Mappable
	if len(key.ReadHashID(assetID).Bytes()) > 0 {
		mappable := assets.mapper.Read(assets.context, key.AssetIDAsKey(assetID))
		if mappable != nil {
			assetList = append(assetList, mappable)
		}
	} else {
		appendMappableList := func(mappable traits.Mappable) bool {
			assetList = append(assetList, mappable)
			return false
		}
		assets.mapper.Iterate(assets.context, key.AssetIDAsKey(assetID), appendMappableList)
	}
	assets.Key, assets.List = key.AssetIDAsKey(assetID), assetList
	return assets
}
func (assets assets) Add(asset mappables.InterNFT) traits.Collection {
	assets.Key = nil
	assets.mapper.Create(assets.context, asset)
	assets.List = append(assets.List, asset)
	return assets
}
func (assets assets) Remove(asset mappables.InterNFT) traits.Collection {
	assets.mapper.Delete(assets.context, key.AssetIDAsKey(asset.GetID()))
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Equal(asset.GetID()) {
			assets.List = append(assets.List[:i], assets.List[i+1:]...)
			break
		}
	}
	return assets
}
func (assets assets) Mutate(asset mappables.InterNFT) traits.Collection {
	assets.mapper.Update(assets.context, asset)
	for i, oldAsset := range assets.List {
		if oldAsset.GetID().Equal(asset.GetID()) {
			assets.List[i] = asset
			break
		}
	}
	return assets
}

func (assets assets) Initialize(context sdkTypes.Context, mapper helpers.Mapper) traits.Collection {
	assets.mapper = mapper
	assets.context = context
	return assets
}
func (assets) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(assets{}, module.Route+"/"+"assets", nil)
}
