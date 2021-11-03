/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Genesis interface {
	Validate() error
	Import(sdkTypes.Context, Mapper, Genesis)
	Export(sdkTypes.Context, Mapper) Genesis

	Encode(codec.JSONMarshaler) []byte
	Decode(codec.JSONMarshaler, []byte) Genesis

	Initialize([]Mappable, []types.Parameter) Genesis

	GetParameters() []types.Parameter
	GetMappableList() []Mappable

	types.Proto
	codecTypes.UnpackInterfacesMessage
}
