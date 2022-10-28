// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Asset)(nil), nil)
	codec.RegisterInterface((*Classification)(nil), nil)
	codec.RegisterInterface((*Document)(nil), nil)
	codec.RegisterInterface((*Identity)(nil), nil)
	codec.RegisterInterface((*Maintainer)(nil), nil)
	codec.RegisterInterface((*Order)(nil), nil)
}
