package mapper

import (
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.InterNFT = (*asset)(nil)

type asset struct {
	AssetID        assetID
	BaseProperties types.BaseProperties
	Lock           types.BaseHeight
	Burn           types.BaseHeight
}

func (asset asset) ID() types.ID { return asset.AssetID }

func (asset asset) ChainID() types.ID          { return asset.AssetID.ChainID }
func (asset asset) ClassificationID() types.ID { return asset.AssetID.ClassificationID }
func (asset asset) MaintainersID() types.ID    { return asset.AssetID.MaintainersID }
func (asset asset) HashID() types.ID           { return asset.AssetID.HashID }

func (asset asset) Properties() types.Properties { return &asset.BaseProperties }

func (asset asset) GetLock() types.Height { return asset.Lock }
func (asset asset) CanSend(currentHeight types.Height) bool {
	return currentHeight.IsGraterThat(asset.Lock)
}

func (asset asset) GetBurn() types.Height { return asset.Burn }
func (asset asset) CanBurn(currentHeight types.Height) bool {
	return currentHeight.IsGraterThat(asset.Burn)
}
func assetFromInterface(interNFT types.InterNFT) asset {
	return asset{
		assetID{
			interNFT.ChainID(),
			interNFT.MaintainersID(),
			interNFT.ClassificationID(),
			interNFT.HashID(),
		},
		types.BasePropertiesFromInterface(interNFT.Properties()),
		types.BaseHeightFromInterface(interNFT.GetLock()),
		types.BaseHeightFromInterface(interNFT.GetBurn()),
	}
}
