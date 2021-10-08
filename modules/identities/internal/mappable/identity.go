/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"strings"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/ids"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	qualifiedMappables "github.com/persistenceOne/persistenceSDK/schema/mappables/qualified"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type identity struct {
	qualifiedMappables.Document //nolint:govet
}

var _ mappables.Identity = (*identity)(nil)

func (identity identity) GetExpiry() types.Property {
	if property := identity.GetProperty(ids.ExpiryProperty); property != nil {
		return property
	}

	return properties.Expiry
}
func (identity identity) GetAuthentication() types.Property {
	if property := identity.GetProperty(ids.AuthenticationProperty); property != nil {
		return property
	}

	return properties.Authentication
}
func (identity identity) GetKey() helpers.Key {
	return key.FromID(identity.ID)
}
func (identity) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, identity{})
}
func (identity identity) IsProvisioned(address sdkTypes.AccAddress) bool {
	if authentication := identity.GetAuthentication(); authentication != nil {
		compareAuthenticationHash := base.NewAccAddressData(address).GenerateHashID().String()

		authenticationHashList := strings.Split(authentication.GetHashID().String(), constants.ListHashStringSeparator)
		for _, authenticationHash := range authenticationHashList {
			if strings.Compare(authenticationHash, compareAuthenticationHash) == 0 {
				return true
			}
		}
	}

	return false
}
func (identity identity) ProvisionAddress(address sdkTypes.AccAddress) mappables.Identity {
	return mappables.Identity(identity)
}

func (identity identity) UnprovisionAddress(address sdkTypes.AccAddress) mappables.Identity {
	return mappables.Identity(identity)
}

func NewIdentity(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Identity {
	return identity{
		Document: qualifiedMappables.Document{
			ID:            id,
			HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties},
			HasMutables:   baseTraits.HasMutables{Properties: mutableProperties},
		},
	}
}
