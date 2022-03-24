/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Classification)(nil), nil)
	codec.RegisterInterface((*Identity)(nil), nil)
	codec.RegisterInterface((*Asset)(nil), nil)
	codec.RegisterInterface((*Maintainer)(nil), nil)
	codec.RegisterInterface((*Meta)(nil), nil)
	codec.RegisterInterface((*Order)(nil), nil)
	codec.RegisterInterface((*Split)(nil), nil)
}
