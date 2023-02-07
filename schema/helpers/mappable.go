// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

type Mappable interface {
	//Size() int
	//Unmarshal([]byte) error
	//MarshalTo([]byte) (int, error)
	GetKey() Key
	RegisterCodec(*codec.LegacyAmino)
	RegisterInterfaces(registry types.InterfaceRegistry)
}
