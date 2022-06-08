// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*DataList)(nil), nil)
	codec.RegisterInterface((*IDList)(nil), nil)
	codec.RegisterInterface((*List)(nil), nil)
	codec.RegisterInterface((*MetaPropertyList)(nil), nil)
	codec.RegisterInterface((*PropertyList)(nil), nil)
	codec.RegisterInterface((*SignatureList)(nil), nil)
}
