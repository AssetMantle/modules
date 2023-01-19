// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/ids"
)

type Split interface {
	GetOwnerID() ids.IdentityID
	GetOwnableID() ids.OwnableID
	GetValue() sdkTypes.Dec
	CanSend(sdkTypes.Dec) bool

	Send(sdkTypes.Dec) Split
	Receive(sdkTypes.Dec) Split
}
