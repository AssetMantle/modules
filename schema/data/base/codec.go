// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, AccAddressData{})
	schema.RegisterModuleConcrete(codec, BooleanData{})
	schema.RegisterModuleConcrete(codec, DecData{})
	schema.RegisterModuleConcrete(codec, HeightData{})
	schema.RegisterModuleConcrete(codec, IDData{})
	schema.RegisterModuleConcrete(codec, StringData{})
	schema.RegisterModuleConcrete(codec, Data_AccAddressData{})
	schema.RegisterModuleConcrete(codec, Data_BooleanData{})
	schema.RegisterModuleConcrete(codec, Data_DecData{})
	schema.RegisterModuleConcrete(codec, Data_HeightData{})
	schema.RegisterModuleConcrete(codec, Data_IdData{})
	schema.RegisterModuleConcrete(codec, Data_StringData{})
	schema.RegisterModuleConcrete(codec, Data{})
	codec.RegisterInterface((*isData_Impl)(nil), nil)
}
