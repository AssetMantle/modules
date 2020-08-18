/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type asset struct {
	ID         types.ID         `json:"id" valid:"required~required field id missing"`
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

func (asset asset) GetImmutables() types.Immutables {
	return asset.Immutables
}

func (asset asset) GetMutables() types.Mutables {
	return asset.Mutables
}

func (asset asset) GetBurn() types.Fact {
	if burnProperty := asset.Immutables.Get().Get(constants.BurnProperty); burnProperty != nil {
		return burnProperty.GetFact()
	} else if burnProperty := asset.Mutables.Get().Get(constants.BurnProperty); burnProperty != nil {
		return burnProperty.GetFact()
	} else {
		return base.NewFact(base.NewHeightData(base.NewHeight(-1)))
	}
}

func (asset asset) GetLock() types.Fact {
	if lockProperty := asset.Immutables.Get().Get(constants.LockProperty); lockProperty != nil {
		return lockProperty.GetFact()
	} else if lockProperty := asset.Mutables.Get().Get(constants.LockProperty); lockProperty != nil {
		return lockProperty.GetFact()
	} else {
		return base.NewFact(base.NewHeightData(base.NewHeight(-1)))
	}
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
func NewAsset(assetID types.ID, immutables types.Immutables, mutables types.Mutables) mappables.InterNFT {
	return asset{
		ID:         assetID,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
