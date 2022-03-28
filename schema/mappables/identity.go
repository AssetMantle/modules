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

type Identity interface {

	// TODO removal for expired identity
	// GetExpiry returns the expiry property of an Identity
	// * If the property is not found, it returns a default value and not nil
	GetExpiry() types.Property

	// GetAuthentication returns the authentication property of an Identity
	// * If the property is not found, it returns a default value and not nil
	GetAuthentication() types.Property

	IsProvisioned(sdkTypes.AccAddress) bool
	ProvisionAddress(sdkTypes.AccAddress) Identity
	UnprovisionAddress(sdkTypes.AccAddress) Identity

	Document
	helpers.Mappable
}
