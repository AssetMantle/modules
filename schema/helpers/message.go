/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Message interface {
	// TODO check if register message code is required
	RegisterLegacyAminoCodec(*codec.LegacyAmino)
	RegisterInterface(registry codecTypes.InterfaceRegistry)
	sdkTypes.Msg
}
