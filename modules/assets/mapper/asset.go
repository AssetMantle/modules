package mapper

import (
	"github.com/persistenceOne/persistenceSDK/types/schema"
)

type asset struct {
	ID         schema.ID
	Burn       schema.Height
	Lock       schema.Height
	Immutables schema.Immutables
	Mutables   schema.Mutables
}

var _ schema.InterNFT = (*asset)(nil)

func (asset asset) GetID() schema.ID {
	return asset.ID
}

func (asset asset) GetChainID() schema.ID {
	return assetIDFromInterface(asset.ID).ChainID
}

func (asset asset) GetClassificationID() schema.ID {
	return assetIDFromInterface(asset.ID).ClassificationID
}

func (asset asset) GetBurn() schema.Height {
	return asset.Burn
}

func (asset asset) CanBurn(currentHeight schema.Height) bool {
	return currentHeight.IsGreaterThan(asset.Burn)
}

func (asset asset) GetLock() schema.Height {
	return asset.Lock
}

func (asset asset) CanSend(currentHeight schema.Height) bool {
	return currentHeight.IsGreaterThan(asset.Lock)
}

func (asset asset) GetImmutables() schema.Immutables {
	return asset.Immutables
}

func (asset asset) GetMutables() schema.Mutables {
	return asset.Mutables
}

func NewAsset(assetID schema.ID, burn schema.Height, lock schema.Height, immutables schema.Immutables, mutables schema.Mutables) schema.InterNFT {
	return asset{
		ID:         assetID,
		Burn:       burn,
		Lock:       lock,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
