// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/identities/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/mappables"
	propertiesSchema "github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type identity struct {
	baseQualified.Document //nolint:govet
}

var _ mappables.Identity = (*identity)(nil)

func (identity identity) GetExpiry() propertiesSchema.Property {
	if property := identity.Document.GetProperty(constants.ExpiryProperty); property != nil {
		return property
	}

	return constants.Expiry
}
func (identity identity) GetAuthentication() propertiesSchema.Property {
	if property := identity.Document.GetProperty(constants.AuthenticationProperty); property != nil {
		return property
	}

	return constants.Authentication
}

// TODO write test cases
func (identity identity) IsProvisioned(accAddress sdkTypes.AccAddress) bool {
	_, found := identity.GetAuthentication().GetDataID().Search(baseData.NewAccAddressData(accAddress).GetID())
	return found
}

// TODO write test cases
func (identity identity) ProvisionAddress(accAddress sdkTypes.AccAddress) mappables.Identity {
	identity.Document = identity.Document.Mutate(
		baseProperties.NewPropertyWithDataID(
			identity.GetAuthentication().GetID(),
			base.NewListDataID(identity.GetAuthentication().GetDataID().Add(baseData.NewAccAddressData(accAddress).GetID())))).(baseQualified.Document)
	return identity
}

// TODO write test cases
func (identity identity) UnprovisionAddress(accAddress sdkTypes.AccAddress) mappables.Identity {
	identity.Document = identity.Document.Mutate(
		baseProperties.NewPropertyWithDataID(
			identity.GetAuthentication().GetID(),
			base.NewListDataID(identity.GetAuthentication().GetDataID().Remove(baseData.NewAccAddressData(accAddress).GetID())))).(baseQualified.Document)
	return identity
}
func (identity identity) GetKey() helpers.Key {
	return key.FromID(identity.Document.ID)
}
func (identity) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, identity{})
}

func NewIdentity(id ids.ID, immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) mappables.Identity {
	return identity{
		Document: baseQualified.Document{
			ID: id,
			// TODO Add classificationID
			Immutables: baseQualified.Immutables{PropertyList: immutableProperties},
			Mutables:   baseQualified.Mutables{Properties: mutableProperties},
		},
	}
}
