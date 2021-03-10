/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type identity struct {
	ID types.ID `json:"id" valid:"required~required field id missing"`
	baseTraits.HasImmutables
	baseTraits.HasMutables //nolint:govet
}

var _ mappables.InterIdentity = (*identity)(nil)

func (identity identity) GetID() types.ID { return identity.ID }
func (identity identity) GetExpiry() types.Property {
	if property := identity.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.Expiry)); property != nil {
		return property
	} else if property := identity.HasMutables.GetMutableProperties().Get(base.NewID(properties.Expiry)); property != nil {
		return property
	} else {
		return base.NewProperty(base.NewID(properties.Expiry), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}
func (identity identity) GetAuthentication() types.Property {
	if property := identity.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.Authentication)); property != nil {
		return property
	} else if property := identity.HasMutables.GetMutableProperties().Get(base.NewID(properties.Authentication)); property != nil {
		return property
	} else {
		return base.NewProperty(base.NewID(properties.Expiry), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}
func (identity identity) GetKey() helpers.Key {
	return key.FromID(identity.ID)
}
func (identity) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, identity{})
}

func NewIdentity(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.InterIdentity {
	return identity{
		ID:            id,
		HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties},
		HasMutables:   baseTraits.HasMutables{Properties: mutableProperties},
	}
}
