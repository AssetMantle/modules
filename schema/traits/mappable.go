/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Mappable interface {
	GetID() types.ID
	Encode() []byte
	Decode([]byte) Mappable
	RegisterCodec(*codec.Codec)
}
