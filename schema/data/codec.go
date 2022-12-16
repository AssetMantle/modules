// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*AccAddressData)(nil), nil)
	codec.RegisterInterface((*BooleanData)(nil), nil)
	codec.RegisterInterface((*Data)(nil), nil)
	codec.RegisterInterface((*DecData)(nil), nil)
	codec.RegisterInterface((*HeightData)(nil), nil)
	codec.RegisterInterface((*IDData)(nil), nil)
	codec.RegisterInterface((*ListData)(nil), nil)
	codec.RegisterInterface((*StringData)(nil), nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface("accAddress", (*AccAddressData)(nil), nil)
	registry.RegisterInterface("booleanData", (*BooleanData)(nil), nil)
	registry.RegisterInterface("Data", (*Data)(nil), nil)
	registry.RegisterInterface("DecData", (*DecData)(nil), nil)
	registry.RegisterInterface("heightData", (*HeightData)(nil), nil)
	registry.RegisterInterface("idData", (*IDData)(nil), nil)
	registry.RegisterInterface("ListData", (*ListData)(nil), nil)
	registry.RegisterInterface("StringData", (*StringData)(nil), nil)
}
