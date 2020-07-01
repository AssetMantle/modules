package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.InterNFTs = (*assets)(nil)

type assets struct {
	ID   types.ID
	List []types.InterNFT

	mapper  Mapper
	context sdkTypes.Context
}

func (Assets assets) GetID() types.ID { return Assets.ID }
func (Assets assets) Get(id types.ID) types.InterNFT {
	for _, oldAsset := range Assets.List {
		if oldAsset.GetID().Compare(id) == 0 {
			return oldAsset
		}
	}
	return nil
}
func (Assets assets) GetList() []types.InterNFT {
	return Assets.List
}

func (Assets assets) Read(id types.ID) types.InterNFTs {
	var assetList []types.InterNFT
	assetsID := assetIDFromInterface(id)
	if assetsID.HashID != nil {
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
	return &assets{id, assetList, Assets.mapper, Assets.context}
}
func (Assets assets) Add(asset types.InterNFT) types.InterNFTs {
	Assets.ID = nil
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

func NewAssets(mapper Mapper, context sdkTypes.Context) types.InterNFTs {
	return &assets{
		ID:      nil,
		List:    nil,
		mapper:  mapper,
		context: context,
	}
}
