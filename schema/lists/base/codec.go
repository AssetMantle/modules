// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema"
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, dataList{})
	schema.RegisterModuleConcrete(codec, idList{})
	schema.RegisterModuleConcrete(codec, list{})
	schema.RegisterModuleConcrete(codec, propertyList{})
	schema.RegisterModuleConcrete(codec, signatureList{})
}
