// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseTypes "github.com/AssetMantle/modules/schema/data/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/constants/properties"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	qualifiedMappables "github.com/AssetMantle/modules/schema/mappables/qualified"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
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
func (identity identity) IsProvisioned(address sdkTypes.AccAddress) bool {
	if authentication := identity.GetAuthentication(); authentication != nil {
		compareAuthenticationHash := baseTypes.NewAccAddressData(address).GenerateHashID().String()

		// TODO impl through list
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
	// TODO
	return nil
}

func (identity identity) UnprovisionAddress(address sdkTypes.AccAddress) mappables.Identity {
	// TODO
	return nil
}
func (identity identity) GetKey() helpers.Key {
	return key.FromID(identity.ID)
}
func (identity) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, identity{})
}

func NewIdentity(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Identity {
	return identity{
		Document: qualifiedMappables.Document{
			ID: id,
			// TODO Add classificationID
			HasImmutables: baseQualified.HasImmutables{Properties: immutableProperties},
			HasMutables:   baseQualified.HasMutables{Properties: mutableProperties},
		},
	}
}
