// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package qualified

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// TODO register package codecs
func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Document)(nil), nil)
	codec.RegisterInterface((*Immutables)(nil), nil)
	codec.RegisterInterface((*Mutables)(nil), nil)
}
