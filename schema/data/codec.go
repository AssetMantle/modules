// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*AccAddressData)(nil), nil)
	codec.RegisterInterface((*BooleanData)(nil), nil)
	codec.RegisterInterface((*Data)(nil), nil)
	codec.RegisterInterface((*DecData)(nil), nil)
	codec.RegisterInterface((*HeightData)(nil), nil)
	codec.RegisterInterface((*IDData)(nil), nil)
	codec.RegisterInterface((*ListData)(nil), nil)
	codec.RegisterInterface((*StringData)(nil), nil)
}
