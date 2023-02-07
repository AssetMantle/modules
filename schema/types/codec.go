// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterInterface((*Height)(nil), nil)
	legacyAmino.RegisterInterface((*Signature)(nil), nil)
	legacyAmino.RegisterInterface((*Split)(nil), nil)
}
