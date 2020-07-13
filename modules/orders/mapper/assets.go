package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.InterNFTs = (*orders)(nil)

type orders struct {
	ID   types.ID
	List []types.InterNFT

	mapper  ordersMapper
	context sdkTypes.Context
}

func (Orders orders) GetID() types.ID { return Orders.ID }
func (Orders orders) Get(id types.ID) types.InterNFT {
	assetID := orderIDFromInterface(id)
	for _, oldAsset := range Orders.List {
		if oldAsset.GetID().Compare(assetID) == 0 {
			return oldAsset
		}
	}
	return nil
}
func (Orders orders) GetList() []types.InterNFT {
	return Orders.List
}

func (Orders orders) Fetch(id types.ID) types.InterNFTs {
	var assetList []types.InterNFT
	assetsID := orderIDFromInterface(id)
	if len(assetsID.HashID.Bytes()) > 0 {
		asset := Orders.mapper.read(Orders.context, assetsID)
		if asset != nil {
			assetList = append(assetList, asset)
		}
	} else {
		appendAssetList := func(asset types.InterNFT) bool {
			assetList = append(assetList, asset)
			return false
		}
		Orders.mapper.iterate(Orders.context, assetsID, appendAssetList)
	}
	return orders{id, assetList, Orders.mapper, Orders.context}
}
func (Orders orders) Add(asset types.InterNFT) types.InterNFTs {
	Orders.ID = nil
	Orders.mapper.create(Orders.context, asset)
	for i, oldAsset := range Orders.List {
		if oldAsset.GetID().Compare(asset.GetID()) < 0 {
			Orders.List = append(append(Orders.List[:i], asset), Orders.List[i+1:]...)
			break
		}
	}
	return Orders
}
func (Orders orders) Remove(asset types.InterNFT) types.InterNFTs {
	Orders.mapper.delete(Orders.context, asset.GetID())
	for i, oldAsset := range Orders.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			Orders.List = append(Orders.List[:i], Orders.List[i+1:]...)
			break
		}
	}
	return Orders
}
func (Orders orders) Mutate(asset types.InterNFT) types.InterNFTs {
	Orders.mapper.update(Orders.context, asset)
	for i, oldAsset := range Orders.List {
		if oldAsset.GetID().Compare(asset.GetID()) == 0 {
			Orders.List[i] = asset
			break
		}
	}
	return Orders
}

func NewAssets(Mapper types.Mapper, context sdkTypes.Context) types.InterNFTs {
	switch mapper := Mapper.(type) {
	case ordersMapper:
		return orders{
			ID:      nil,
			List:    nil,
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", constants.ModuleName)))
	}

}
