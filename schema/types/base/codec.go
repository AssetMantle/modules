// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

const moduleName = "types"

// TODO pick module name programmatically
func RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, moduleName, height{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, parameter{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, property{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, signature{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, signatures{})
}
