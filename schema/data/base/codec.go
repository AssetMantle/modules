// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, accAddressData{})
	schema.RegisterModuleConcrete(codec, booleanData{})
	schema.RegisterModuleConcrete(codec, decData{})
	schema.RegisterModuleConcrete(codec, heightData{})
	schema.RegisterModuleConcrete(codec, IDData{})
	schema.RegisterModuleConcrete(codec, listData{})
	schema.RegisterModuleConcrete(codec, stringData{})
}
