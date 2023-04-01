// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/types"
)

type Identity interface {

	// TODO removal for expired identity
	// GetExpiry returns the expiry property of an Identity
	// * If the property is not found, it returns a default value and not nil
	GetExpiry() types.Height

	// GetAuthentication returns the authentication property of an Identity
	// * If the property is not found, it returns a default value and not nil
	GetAuthentication() data.ListData

	IsProvisioned(sdkTypes.AccAddress) bool
	GetProvisionedAddressCount() int64
	ProvisionAddress(...sdkTypes.AccAddress) Identity
	UnprovisionAddress(...sdkTypes.AccAddress) Identity

	Document
}
