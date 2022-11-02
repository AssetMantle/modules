// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, asset{})
	codecUtilities.RegisterModuleConcrete(codec, classification{})
	codecUtilities.RegisterModuleConcrete(codec, document{})
	codecUtilities.RegisterModuleConcrete(codec, identity{})
	codecUtilities.RegisterModuleConcrete(codec, maintainer{})
	codecUtilities.RegisterModuleConcrete(codec, order{})
}
