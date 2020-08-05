package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Asset struct {
	ID         types.ID         `json:"id" valid:"required~required field id missing"`
	Burn       types.Height     `json:"burn" valid:"required~required field burn missing"`
	Lock       types.Height     `json:"lock" valid:"required field lock missing"`
	Immutables types.Immutables `json:"immutables" valid:"required field immutables missing"`
	Mutables   types.Mutables   `json:"mutables" valid:"required~required field mutables missing"`
}

var _ mappables.InterNFT = (*Asset)(nil)

func (asset Asset) GetID() types.ID {
	return asset.ID
}

func (asset Asset) GetChainID() types.ID {
	return assetIDFromInterface(asset.ID).ChainID
}

func (asset Asset) GetClassificationID() types.ID {
	return assetIDFromInterface(asset.ID).ClassificationID
}

func (asset Asset) GetBurn() types.Height {
	return asset.Burn
}

func (asset Asset) CanBurn(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(asset.Burn)
}

func (asset Asset) GetLock() types.Height {
	return asset.Lock
}

func (asset Asset) CanSend(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(asset.Lock)
}

func (asset Asset) GetImmutables() types.Immutables {
	return asset.Immutables
}

func (asset Asset) GetMutables() types.Mutables {
	return asset.Mutables
}
func (asset Asset) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(asset)
}
func (asset Asset) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &asset)
	return asset
}
func assetPrototype() traits.Mappable {
	return Asset{}
}
func NewAsset(assetID types.ID, burn types.Height, lock types.Height, immutables types.Immutables, mutables types.Mutables) mappables.InterNFT {
	return Asset{
		ID:         assetID,
		Burn:       burn,
		Lock:       lock,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
