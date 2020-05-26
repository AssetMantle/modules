package mapper

import "github.com/persistenceOne/persistenceSDK/types"

func baseAssetFromInterface(asset types.Asset) baseAsset {
	return baseAsset{
		baseAssetID{
			asset.ChainID(),
			asset.MaintainersID(),
			asset.ClassificationID(),
			asset.HashID(),
		},
		asset.OwnersID(),
		asset.Properties(),
		asset.GetLock(),
		asset.GetBurn(),
	}
}
