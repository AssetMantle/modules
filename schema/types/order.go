// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
)

type Order interface {
	GetExchangeRate() sdkTypes.Dec
	GetCreationHeight() Height
	GetMakerOwnableID() ids.OwnableID
	GetTakerOwnableID() ids.OwnableID
	GetMakerID() ids.IdentityID
	GetTakerID() ids.IdentityID
	GetExpiryHeight() Height
	GetMakerOwnableSplit() sdkTypes.Dec

	qualified.Document
}
