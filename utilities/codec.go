// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"reflect"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
)

func MakeModuleCode(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()
	keyPrototype().RegisterCodec(Codec)
	mappablePrototype().RegisterCodec(Codec)
	schema.RegisterCodec(Codec)
	Codec.Seal()

	return Codec
}

func RegisterModuleConcrete(codec *codec.LegacyAmino, o interface{}) {
	codec.RegisterConcrete(o, reflect.TypeOf(o).PkgPath()+"/"+reflect.TypeOf(o).Name(), nil)
}

func MakeMessageCodec(messagePrototype func() helpers.Message) *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()
	messagePrototype().RegisterCodec(Codec)
	schema.RegisterCodec(Codec)
	Codec.Seal()

	return Codec
}
