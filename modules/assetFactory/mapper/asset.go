package mapper

import (
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.InterNFT = (*asset)(nil)

type asset struct {
	ID         types.ID
	Properties types.Properties
	Lock       types.Height
	Burn       types.Height
}

func (asset asset) GetID() types.ID { return asset.ID }

func (asset asset) GetChainID() types.ID { return assetIDFromInterface(asset.ID).ChainID }
func (asset asset) GetClassificationID() types.ID {
	return assetIDFromInterface(asset.ID).ClassificationID
}
func (asset asset) GetMaintainersID() types.ID { return assetIDFromInterface(asset.ID).MaintainersID }
func (asset asset) GetHashID() types.ID        { return assetIDFromInterface(asset.ID).HashID }

func (asset asset) GetProperties() types.Properties { return asset.Properties }

func (asset asset) GetLock() types.Height { return asset.Lock }
func (asset asset) CanSend(currentHeight types.Height) bool {
	return currentHeight.IsGraterThat(asset.Lock)
}

func (asset asset) GetBurn() types.Height { return asset.Burn }
func (asset asset) CanBurn(currentHeight types.Height) bool {
	return currentHeight.IsGraterThat(asset.Burn)
}
func NewAsset(assetID types.ID, properties types.Properties, lock types.Height, burn types.Height) types.InterNFT {
	return &asset{
		ID:         assetID,
		Properties: properties,
		Lock:       lock,
		Burn:       burn,
	}
}
