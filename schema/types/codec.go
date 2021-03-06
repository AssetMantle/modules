/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*Data)(nil), nil)
	codec.RegisterInterface((*Height)(nil), nil)
	codec.RegisterInterface((*ID)(nil), nil)
	codec.RegisterInterface((*MetaProperties)(nil), nil)
	codec.RegisterInterface((*MetaProperty)(nil), nil)
	codec.RegisterInterface((*Parameter)(nil), nil)
	codec.RegisterInterface((*Properties)(nil), nil)
	codec.RegisterInterface((*Property)(nil), nil)
	codec.RegisterInterface((*Signature)(nil), nil)
	codec.RegisterInterface((*Signatures)(nil), nil)
}
