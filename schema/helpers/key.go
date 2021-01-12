/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import "github.com/cosmos/cosmos-sdk/codec"

type Key interface {
	GenerateStoreKeyBytes() []byte
	RegisterCodec(*codec.Codec)
	IsPartial() bool
	Matches(Key) bool
}
