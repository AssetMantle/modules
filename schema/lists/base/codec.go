// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

const moduleName = "lists"

func RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, moduleName, dataList{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, idList{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, list{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, metaPropertyList{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, propertyList{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, signatureList{})
}
