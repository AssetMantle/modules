// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/constants/properties"
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type asset struct {
	base.Document //nolint:govet
}

var _ mappables.Asset = (*asset)(nil)

func (asset asset) GetBurn() types.Property {
	if burn := asset.GetProperty(ids.BurnProperty); burn != nil {
		return burn
	}

	return properties.Burn
}
func (asset asset) GetLock() types.Property {
	if lock := asset.GetProperty(ids.LockProperty); lock != nil {
		return lock
	}

	return properties.Lock
}

// TODO change to supply
func (asset asset) GetValue() types.Property {
	if value := asset.GetProperty(ids.ValueProperty); value != nil {
		return value
	}

	return properties.Value
}
func (asset asset) GetKey() helpers.Key {
	return key.FromID(asset.ID)
}
func (asset) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, asset{})
}

func NewAsset(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Asset {
	return asset{
		Document: base.Document{
			ID:               id,
			ClassificationID: key.ReadClassificationID(id),
			HasImmutables:    base.HasImmutables{Properties: immutableProperties},
			HasMutables:      base.HasMutables{Properties: mutableProperties},
		},
	}
}
