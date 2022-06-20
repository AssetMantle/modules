// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Error)(nil), nil)
	codec.RegisterInterface((*Mappable)(nil), nil)
	codec.RegisterInterface((*QueryRequest)(nil), nil)
	codec.RegisterInterface((*QueryResponse)(nil), nil)
	codec.RegisterInterface((*TransactionRequest)(nil), nil)
}
