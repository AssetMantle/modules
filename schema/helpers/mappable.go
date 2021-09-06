/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Mappable interface {
	GetKey() Key
	RegisterLegacyAminoCodec(amino *codec.LegacyAmino)

	types.Proto
	Reset()
	String() string
	ProtoMessage()
	Marshal() ([]byte, error)
	MarshalToSizedBuffer([]byte) (int, error)
}
