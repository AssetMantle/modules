// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterInterface((*Asset)(nil), nil)
	legacyAmino.RegisterInterface((*Classification)(nil), nil)
	legacyAmino.RegisterInterface((*Document)(nil), nil)
	legacyAmino.RegisterInterface((*Identity)(nil), nil)
	legacyAmino.RegisterInterface((*Maintainer)(nil), nil)
	legacyAmino.RegisterInterface((*Order)(nil), nil)
}
