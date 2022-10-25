// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Asset)(nil), nil)
	codec.RegisterInterface((*Classification)(nil), nil)
	codec.RegisterInterface((*Height)(nil), nil)
	codec.RegisterInterface((*Identity)(nil), nil)
	codec.RegisterInterface((*Maintainer)(nil), nil)
	codec.RegisterInterface((*Meta)(nil), nil)
	codec.RegisterInterface((*Order)(nil), nil)
	codec.RegisterInterface((*Signature)(nil), nil)
	codec.RegisterInterface((*Split)(nil), nil)
}
