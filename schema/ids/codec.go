// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*DataID)(nil), nil)
	codec.RegisterInterface((*ID)(nil), nil)
	codec.RegisterInterface((*PropertyID)(nil), nil)
}
