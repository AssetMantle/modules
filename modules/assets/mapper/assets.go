package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.InterNFTs = (*assets)(nil)

type assets struct {
	ID   types.ID
	List []types.InterNFT

	Mapper  Mapper
	Context sdkTypes.Context
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
		assetList = append(assetList, Assets.Mapper.read(Assets.Context, assetsID))
	} else {
		appendAssetList := func(asset types.InterNFT) bool {
			assetList = append(assetList, asset)
			return false
		}
		Assets.Mapper.iterate(Assets.Context, assetsID, appendAssetList)
	}
	return &assets{id, assetList, Assets.Mapper, Assets.Context}
}
func (Assets assets) Add(asset types.InterNFT) types.InterNFTs {
	Assets.ID = nil
	Assets.Mapper.create(Assets.Context, asset)
	for i, oldAsset := range Assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) < 0 {
			Assets.List = append(append(Assets.List[:i], asset), Assets.List[i+1:]...)
			break
		}
	}
	return Assets
}
func (Assets assets) Remove(asset types.InterNFT) types.InterNFTs {
	Assets.Mapper.delete(Assets.Context, asset.GetID())
	for i, oldAsset := range Assets.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			Assets.List = append(Assets.List[:i], Assets.List[i+1:]...)
			break
		}
	}
	return Assets
}
func (Assets assets) Mutate(asset types.InterNFT) types.InterNFTs {
	Assets.Mapper.update(Assets.Context, asset)
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
		Mapper:  mapper,
		Context: context,
	}
}
