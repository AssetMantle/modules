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
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/mappables"
	properties2 "github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type asset struct {
	base.Document //nolint:govet
}

var _ mappables.Asset = (*asset)(nil)

func (asset asset) GetBurn() properties2.Property {
	if burn := asset.GetProperty(ids.BurnProperty); burn != nil {
		return burn
	}

	return properties.Burn
}
func (asset asset) GetLock() properties2.Property {
	if lock := asset.GetProperty(ids.LockProperty); lock != nil {
		return lock
	}

	return properties.Lock
}
func (asset asset) GetSupply() properties2.Property {
	if supply := asset.GetProperty(ids.SupplyProperty); supply != nil {
		return supply
	}

	return properties.Supply
}
func (asset asset) GetKey() helpers.Key {
	return key.FromID(asset.ID)
}
func (asset) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, asset{})
}

func NewAsset(id ids2.ID, immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) mappables.Asset {
	return asset{
		Document: base.Document{
			ID:               id,
			ClassificationID: key.ReadClassificationID(id),
			Immutables:       base.Immutables{PropertyList: immutableProperties},
			Mutables:         base.Mutables{Properties: mutableProperties},
		},
	}
}
