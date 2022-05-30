// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	constantIDs "github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/constants/properties"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/mappables"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type identity struct {
	baseQualified.Document //nolint:govet
}

var _ mappables.Identity = (*identity)(nil)

func (identity identity) GetExpiry() types.Property {
	if property := identity.Document.GetProperty(constantIDs.ExpiryProperty); property != nil {
		return property
	}

	return properties.Expiry
}
func (identity identity) GetAuthentication() types.Property {
	if property := identity.Document.GetProperty(constantIDs.AuthenticationProperty); property != nil {
		return property
	}

	return properties.Authentication
}

// TODO write test cases
func (identity identity) IsProvisioned(accAddress sdkTypes.AccAddress) bool {
	_, found := identity.GetAuthentication().GetDataID().(ids.ListDataID).Search(baseData.NewAccAddressData(accAddress).GetID())
	return found
}

// TODO write test cases
func (identity identity) ProvisionAddress(accAddress sdkTypes.AccAddress) mappables.Identity {
	identity.Document = identity.Document.Mutate(
		baseTypes.NewPropertyWithDataID(
			identity.GetAuthentication().GetID(),
			base.NewListID(identity.GetAuthentication().GetDataID().(ids.ListDataID).Add(baseData.NewAccAddressData(accAddress).GetID())))).(baseQualified.Document)
	return identity
}

// TODO write test cases
func (identity identity) UnprovisionAddress(accAddress sdkTypes.AccAddress) mappables.Identity {
	identity.Document = identity.Document.Mutate(
		baseTypes.NewPropertyWithDataID(
			identity.GetAuthentication().GetID(),
			base.NewListID(identity.GetAuthentication().GetDataID().(ids.ListDataID).Remove(baseData.NewAccAddressData(accAddress).GetID())))).(baseQualified.Document)
	return identity
}
func (identity identity) GetKey() helpers.Key {
	return key.FromID(identity.Document.ID)
}
func (identity) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, identity{})
}

func NewIdentity(id types.ID, immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) mappables.Identity {
	return identity{
		Document: baseQualified.Document{
			ID: id,
			// TODO Add classificationID
			Immutables: baseQualified.Immutables{PropertyList: immutableProperties},
			Mutables:   baseQualified.Mutables{Properties: mutableProperties},
		},
	}
}
