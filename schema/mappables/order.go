// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappables

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
)

type Order interface {
	GetExchangeRate() sdkTypes.Dec
	GetCreationHeight() types.Height
	GetMakerOwnableID() ids.OwnableID
	GetTakerOwnableID() ids.OwnableID
	GetMakerID() ids.IdentityID
	GetTakerID() ids.IdentityID
	GetExpiryHeight() types.Height
	GetMakerOwnableSplit() sdkTypes.Dec

	qualified.Document
	helpers.Mappable
}
