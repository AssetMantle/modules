// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	schemaCodec "github.com/AssetMantle/schema/go/codec"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/helpers"
)

func RegisterLegacyAminoCodec(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()
	keyPrototype().RegisterLegacyAminoCodec(Codec)
	mappablePrototype().RegisterLegacyAminoCodec(Codec)
	schemaCodec.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()

	return Codec
}
