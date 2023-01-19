// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterInterface((*AccAddressData)(nil), nil)
	legacyAmino.RegisterInterface((*AnyData)(nil), nil)
	legacyAmino.RegisterInterface((*BooleanData)(nil), nil)
	legacyAmino.RegisterInterface((*Data)(nil), nil)
	legacyAmino.RegisterInterface((*DecData)(nil), nil)
	legacyAmino.RegisterInterface((*HeightData)(nil), nil)
	legacyAmino.RegisterInterface((*IDData)(nil), nil)
	legacyAmino.RegisterInterface((*ListData)(nil), nil)
	legacyAmino.RegisterInterface((*StringData)(nil), nil)
}
