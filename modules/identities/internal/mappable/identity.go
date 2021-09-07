/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/ids"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappables/qualified"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type identity struct {
	qualified.Document
}

var _ mappables.Identity = (*identity)(nil)

func (identity identity) GetExpiry() types.Property {
	if property := identity.GetProperty(ids.Expiry); property != nil {
		return property
	}

	return properties.Expiry
}
func (identity identity) GetAuthentication() types.Property {
	if property := identity.GetProperty(ids.Authentication); property != nil {
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
	//flag := false
	//accAddressListData, ok := identity.GetAuthentication().GetFact().(types.ListData)
	//
	//if !ok {
	//	panic(errors.IncorrectFormat)
	//}
	//
	//if !address.Empty() && accAddressListData.Search(qualified.NewAccAddressData(address)) != -1 {
	//	flag = true
	//}
	impl
	return true
}
func (identity identity) IsUnprovisioned(address sdkTypes.AccAddress) bool {
	//flag := false
	//accAddressListData, ok := identity.GetAuthentication().GetFact().(types.ListData)
	//
	//if !ok {
	//	panic(errors.IncorrectFormat)
	//}
	//
	//if !address.Empty() && !(accAddressListData.Search(qualified.NewAccAddressData(address)) != -1) {
	//	flag = true
	//}
	impl
	return true
}
func (identity identity) ProvisionAddress(address sdkTypes.AccAddress) mappables.Identity {
	impl
	return mappables.Identity(identity)
}
func (identity identity) UnprovisionAddress(address sdkTypes.AccAddress) mappables.Identity {
	impl
	return mappables.Identity(identity)
}

func NewIdentity(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Identity {
	return identity{
		Document: qualified.Document{
			ID:            id,
			HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties},
			HasMutables:   baseTraits.HasMutables{Properties: mutableProperties},
		},
	}
}
