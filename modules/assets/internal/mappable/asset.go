/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
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
		data, _ := base.ReadHeightData("")
		return base.NewProperty(base.NewID(properties.Burn), base.NewFact(data))
	}
}

func (asset asset) GetLock() types.Property {
	if lockProperty := asset.Immutables.Get().Get(base.NewID(properties.Lock)); lockProperty != nil {
		return lockProperty
	} else if lockProperty := asset.Mutables.Get().Get(base.NewID(properties.Lock)); lockProperty != nil {
		return lockProperty
	} else {
		data, _ := base.ReadHeightData("")
		return base.NewProperty(base.NewID(properties.Lock), base.NewFact(data))
	}
}
func (asset asset) GetKey() helpers.Key {
	return key.New(asset.ID)
}
func (asset) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(asset{}, constants.ProjectRoute+"/"+"asset", nil)
}
func NewAsset(assetID types.ID, immutables types.Immutables, mutables types.Mutables) mappables.InterNFT {
	return asset{
		ID:         assetID,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
