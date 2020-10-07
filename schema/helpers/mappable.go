/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

type Mappable interface {
	GetKey() Key
	Encode() []byte
	Decode([]byte) Mappable
	RegisterCodec(*codec.Codec)
}
