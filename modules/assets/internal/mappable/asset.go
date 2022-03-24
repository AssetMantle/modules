/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants/ids" //nolint:typecheck
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type asset struct {
	qualified.Document //nolint:govet
}

var _ mappables.Asset = (*asset)(nil) //nolint:typecheck

func (asset asset) GetBurn() types.Property {
	if burn := asset.GetProperty(ids.BurnProperty); burn != nil { //nolint:typecheck
		return burn
	}

	return properties.Burn
}
func (asset asset) GetLock() types.Property {
	if lock := asset.GetProperty(ids.LockProperty); lock != nil { //nolint:typecheck
		return lock
	}

	return properties.Lock
}
func (asset asset) GetValue() types.Property {
	if value := asset.GetProperty(ids.ValueProperty); value != nil { //nolint:typecheck
		return value
	}

	return properties.Value
}
func (asset asset) GetKey() helpers.Key {
	return key.FromID(asset.ID) //nolint:typecheck
}
func (asset) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, asset{})
}

func NewAsset(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Asset {
	return asset{ //nolint:typecheck
		Document: qualified.Document{
			ID:               id,
			ClassificationID: key.ReadClassificationID(id),
			HasImmutables:    qualified.HasImmutables{Properties: immutableProperties},
			HasMutables:      qualified.HasMutables{Properties: mutableProperties},
		},
	}
}
