/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"github.com/persistenceOne/persistenceSDK/modules/classifications/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type asset struct {
	ID         types.ID         `json:"id" valid:"required~required field id missing"`
	Burn       types.Height     `json:"burn" valid:"required~required field burn missing"`
	Lock       types.Height     `json:"lock" valid:"required field lock missing"`
	Immutables types.Immutables `json:"immutables" valid:"required field immutables missing"`
	Mutables   types.Mutables   `json:"mutables" valid:"required~required field mutables missing"`
}

var _ mappables.InterNFT = (*asset)(nil)

func (asset asset) GetID() types.ID {
	return asset.ID
}

func (asset asset) GetChainID() types.ID {
	return mapper.ChainIDFromClassificationID(assetIDFromInterface(asset.ID).ClassificationID)
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
func (asset asset) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(asset)
}
func (asset asset) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &asset)
	return asset
}
func assetPrototype() traits.Mappable {
	return asset{}
}
func NewAsset(assetID types.ID, burn types.Height, lock types.Height, immutables types.Immutables, mutables types.Mutables) mappables.InterNFT {
	return asset{
		ID:         assetID,
		Burn:       burn,
		Lock:       lock,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
