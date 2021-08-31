/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Identity interface {
	GetExpiry() types.Property
	GetAuthentication() types.Property

	IsProvisioned(sdkTypes.AccAddress) bool
	IsUnprovisioned(sdkTypes.AccAddress) bool
	ProvisionAddress(sdkTypes.AccAddress) Identity
	UnprovisionAddress(sdkTypes.AccAddress) Identity

	Document
}
