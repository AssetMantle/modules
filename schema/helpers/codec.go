/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Mappable)(nil), nil)
	codec.RegisterInterface((*QueryRequest)(nil), nil)
	codec.RegisterInterface((*QueryResponse)(nil), nil)
	codec.RegisterInterface((*TransactionRequest)(nil), nil)
}
