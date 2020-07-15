package mapper

import (
	"github.com/persistenceOne/persistenceSDK/types"
)

type asset struct {
	ID         types.ID
	Burn       types.Height
	Lock       types.Height
	Immutables types.Immutables
	Mutables   types.Mutables
}

var _ types.InterNFT = (*asset)(nil)

func (asset asset) GetID() types.ID {
	return asset.ID
}

func (asset asset) GetChainID() types.ID {
	return assetIDFromInterface(asset.ID).ChainID
}

func (asset asset) GetClassificationID() types.ID {
	return assetIDFromInterface(asset.ID).ClassificationID
}

func (asset asset) GetBurn() types.Height {
	return asset.Burn
}

func (asset asset) CanBurn(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(asset.Burn)
}

func (asset asset) GetLock() types.Height {
	return asset.Lock
}

func (asset asset) CanSend(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(asset.Lock)
}

func (asset asset) GetImmutables() types.Immutables {
	return asset.Immutables
}

func (asset asset) GetMutables() types.Mutables {
	return asset.Mutables
}

func NewAsset(assetID types.ID, burn types.Height, lock types.Height, immutables types.Immutables, mutables types.Mutables) types.InterNFT {
	return asset{
		ID:         assetID,
		Burn:       burn,
		Lock:       lock,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
