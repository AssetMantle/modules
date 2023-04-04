// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	schema "github.com/AssetMantle/schema/x"
	"github.com/AssetMantle/schema/x/helpers"
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()
	keyPrototype().RegisterLegacyAminoCodec(Codec)
	mappablePrototype().RegisterLegacyAminoCodec(Codec)
	schema.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()

	return Codec
}
