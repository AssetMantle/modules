// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterInterface((*IDList)(nil), nil)
	legacyAmino.RegisterInterface((*List)(nil), nil)
	legacyAmino.RegisterInterface((*PropertyList)(nil), nil)
	legacyAmino.RegisterInterface((*SignatureList)(nil), nil)
}
