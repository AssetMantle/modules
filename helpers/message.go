// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Message interface {
	// TODO check if register message code is required
	RegisterLegacyAminoCodec(*codec.LegacyAmino)
	RegisterInterface(types.InterfaceRegistry)
	Type() string
	sdkTypes.Msg
}
