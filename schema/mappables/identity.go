// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappables

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
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

	qualified.Document
	helpers.Mappable
}
