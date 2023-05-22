// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transaction

import (
	"github.com/cosmos/cosmos-sdk/codec"

	schemaCodec "github.com/AssetMantle/schema/go/codec"

	"github.com/AssetMantle/modules/helpers"
)

func RegisterLegacyAminoCodec(messagePrototype func() helpers.Message) *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()
	messagePrototype().RegisterLegacyAminoCodec(Codec)
	schemaCodec.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()

	return Codec
}
