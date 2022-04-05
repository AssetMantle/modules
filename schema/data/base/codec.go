// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

const moduleName = "types"

func RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, moduleName, accAddressData{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, booleanData{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, listData{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, decData{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, heightData{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, idData{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, stringData{})
}
