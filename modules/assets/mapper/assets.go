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

func (Assets assets) GetID() types.ID { return Assets.ID }
func (Assets assets) Get(id types.ID) types.InterNFT {
	assetID := assetIDFromInterface(id)
	for _, oldAsset := range Assets.List {
		if oldAsset.GetID().Compare(assetID) == 0 {
			return oldAsset
		}
	}
	return nil
}
func (Assets assets) GetList() []types.InterNFT {
	return Assets.List
}

func (Assets assets) Fetch(id types.ID) types.InterNFTs {
	var assetList []types.InterNFT
	assetsID := assetIDFromInterface(id)
	if len(assetsID.HashID.Bytes()) > 0 {
		asset := Assets.mapper.read(Assets.context, assetsID)
		if asset != nil {
			assetList = append(assetList, asset)
		}
	} else {
		appendAssetList := func(asset types.InterNFT) bool {
			assetList = append(assetList, asset)
			return false
		}
		Assets.mapper.iterate(Assets.context, assetsID, appendAssetList)
	}
	return assets{id, assetList, Assets.mapper, Assets.context}
}
func (Assets assets) Add(asset types.InterNFT) types.InterNFTs {
	Assets.ID = readAssetID("")
	Assets.mapper.create(Assets.context, asset)
	for i, oldAsset := range Assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) < 0 {
			Assets.List = append(append(Assets.List[:i], asset), Assets.List[i+1:]...)
			break
		}
	}
	return Assets
}
func (Assets assets) Remove(asset types.InterNFT) types.InterNFTs {
	Assets.mapper.delete(Assets.context, asset.GetID())
	for i, oldAsset := range Assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			Assets.List = append(Assets.List[:i], Assets.List[i+1:]...)
			break
		}
	}
	return Assets
}
func (Assets assets) Mutate(asset types.InterNFT) types.InterNFTs {
	Assets.mapper.update(Assets.context, asset)
	for i, oldAsset := range Assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			Assets.List[i] = asset
			break
		}
	}
	return Assets
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
