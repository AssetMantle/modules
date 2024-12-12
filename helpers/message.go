// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/AssetMantle/schema/ids"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Message interface {
	GetFromAddress() sdkTypes.AccAddress
	GetFromIdentityID() ids.IdentityID
	RegisterInterface(types.InterfaceRegistry)

	sdkTypes.Msg
}
