/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
)

type InterIdentity interface {
	GetProvisionedAddressList() []sdkTypes.AccAddress
	GetUnprovisionedAddressList() []sdkTypes.AccAddress

	ProvisionAddress(sdkTypes.AccAddress) InterIdentity
	UnprovisionAddress(sdkTypes.AccAddress) InterIdentity

	IsProvisioned(sdkTypes.AccAddress) bool
	IsUnprovisioned(sdkTypes.AccAddress) bool

	traits.HasImmutables
	traits.HasMutables
	helpers.Mappable
}
