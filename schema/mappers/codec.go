/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappers

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Chains)(nil), nil)
	codec.RegisterInterface((*InterIdentities)(nil), nil)
	codec.RegisterInterface((*InterNFTs)(nil), nil)
	codec.RegisterInterface((*Maintainers)(nil), nil)
	codec.RegisterInterface((*Metas)(nil), nil)
	codec.RegisterInterface((*Orders)(nil), nil)
	codec.RegisterInterface((*Splits)(nil), nil)
}
