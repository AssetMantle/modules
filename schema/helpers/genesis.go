// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"

	"github.com/AssetMantle/modules/schema/parameters"
)

type Genesis interface {
	proto.Message

	Default() Genesis
	Validate() error
	Import(sdkTypes.Context, Mapper, Parameters)
	Export(sdkTypes.Context, Mapper, Parameters) Genesis

	Encode(codec.JSONCodec) []byte
	Decode(codec.JSONCodec, []byte) Genesis

	Initialize([]Mappable, []parameters.Parameter) Genesis

	GetParameterList() []parameters.Parameter
	GetMappableList() []Mappable
}
