// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	baseTraits "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*error)(nil), nil)
	types.RegisterCodec(codec)
	baseTypes.RegisterCodec(codec)
	traits.RegisterCodec(codec)
	baseTraits.RegisterCodec(codec)
	mappables.RegisterCodec(codec)
	helpers.RegisterCodec(codec)
}
