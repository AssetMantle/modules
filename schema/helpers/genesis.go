// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"

	sdkCodec "github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/parameters"
)

type Genesis interface {
	Default() Genesis
	Validate() error
	Import(context.Context, Mapper, Parameters)
	Export(context.Context, Mapper, Parameters) Genesis

	Encode(sdkCodec.JSONCodec) []byte
	Decode(sdkCodec.JSONCodec, []byte) Genesis

	Initialize([]Mappable, []parameters.Parameter) Genesis

	GetParameterList() []parameters.Parameter
	GetMappableList() []Mappable
}
