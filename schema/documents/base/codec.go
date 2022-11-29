// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, asset{})
	schema.RegisterModuleConcrete(codec, classification{})
	schema.RegisterModuleConcrete(codec, document{})
	schema.RegisterModuleConcrete(codec, identity{})
	schema.RegisterModuleConcrete(codec, maintainer{})
	schema.RegisterModuleConcrete(codec, order{})
}
