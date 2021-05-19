/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterIdentity interface {
	GetAuthentication() types.Property
	GetExpiry() types.Property
	IsProvisioned(address sdkTypes.AccAddress) bool
	IsUnprovisioned(address sdkTypes.AccAddress) bool
	UnprovisionAddress(address sdkTypes.AccAddress) helpers.Mappable
	Document
}
