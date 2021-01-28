/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
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
func (asset asset) GetClassificationID() types.ID {
	return key.ReadClassificationID(asset.ID)
}
func (asset asset) GetImmutables() types.Immutables {
	return asset.Immutables
}
func (asset asset) GetMutables() types.Mutables {
	return asset.Mutables
}
func (asset asset) GetBurn() types.Property {
	if burnProperty := asset.Immutables.Get().Get(base.NewID(properties.Burn)); burnProperty != nil {
		return burnProperty
	} else if burnProperty := asset.Mutables.Get().Get(base.NewID(properties.Burn)); burnProperty != nil {
		return burnProperty
	} else {
		return base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}
func (asset asset) GetLock() types.Property {
	if lockProperty := asset.Immutables.Get().Get(base.NewID(properties.Lock)); lockProperty != nil {
		return lockProperty
	} else if lockProperty := asset.Mutables.Get().Get(base.NewID(properties.Lock)); lockProperty != nil {
		return lockProperty
	} else {
		return base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}
func (asset asset) GetValue() types.Property {
	if splitProperty := asset.Immutables.Get().Get(base.NewID(properties.Value)); splitProperty != nil {
		return splitProperty
	} else if splitProperty := asset.Mutables.Get().Get(base.NewID(properties.Value)); splitProperty != nil {
		return splitProperty
	} else {
		return base.NewProperty(base.NewID(properties.Value), base.NewFact(base.NewDecData(sdkTypes.SmallestDec())))
	}
}
func (asset asset) GetKey() helpers.Key {
	return key.New(asset.ID)
}
func (asset) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, asset{})
}

func NewAsset(assetID types.ID, immutables types.Immutables, mutables types.Mutables) mappables.InterNFT {
	return asset{
		ID:         assetID,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
