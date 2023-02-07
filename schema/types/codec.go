// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterCodec(codec *codec.Codec) {
	types.RegisterCodec(codec)
	codec.RegisterInterface((*Height)(nil), nil)
	codec.RegisterInterface((*Signature)(nil), nil)
	codec.RegisterInterface((*Split)(nil), nil)
}
