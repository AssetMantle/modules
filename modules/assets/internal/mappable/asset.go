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
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type asset struct {
	ID types.ID `json:"id" valid:"required~required field id missing"`
	baseTraits.Immutables
	baseTraits.Mutables
}

var _ mappables.InterNFT = (*asset)(nil)

func (asset asset) GetID() types.ID {
	return asset.ID
}
func (asset asset) GetClassificationID() types.ID {
	return key.ReadClassificationID(asset.ID)
}
func (asset asset) GetBurn() types.Property {
	if burnProperty := asset.Immutables.GetImmutableProperties().Get(base.NewID(properties.Burn)); burnProperty != nil {
		return burnProperty
	} else if burnProperty := asset.Mutables.GetMutableProperties().Get(base.NewID(properties.Burn)); burnProperty != nil {
		return burnProperty
	} else {
		return base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}
func (asset asset) GetLock() types.Property {
	if lockProperty := asset.Immutables.GetImmutableProperties().Get(base.NewID(properties.Lock)); lockProperty != nil {
		return lockProperty
	} else if lockProperty := asset.Mutables.GetMutableProperties().Get(base.NewID(properties.Lock)); lockProperty != nil {
		return lockProperty
	} else {
		return base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}
func (asset asset) GetValue() types.Property {
	if splitProperty := asset.Immutables.GetImmutableProperties().Get(base.NewID(properties.Value)); splitProperty != nil {
		return splitProperty
	} else if splitProperty := asset.Mutables.GetMutableProperties().Get(base.NewID(properties.Value)); splitProperty != nil {
		return splitProperty
	} else {
		return base.NewProperty(base.NewID(properties.Value), base.NewFact(base.NewDecData(sdkTypes.SmallestDec())))
	}
}
func (asset asset) GetKey() helpers.Key {
	return key.FromID(asset.ID)
}
func (asset) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, asset{})
}

func NewAsset(assetID types.ID, immutables types.Properties, mutables types.Properties) mappables.InterNFT {
	return asset{
		ID:         assetID,
		Immutables: baseTraits.Immutables{Properties: immutables},
		Mutables:   baseTraits.Mutables{Properties: mutables},
	}
}
