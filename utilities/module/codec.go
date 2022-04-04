// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
)

func RegisterCodec(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) *codec.Codec {
	Codec := codec.New()
	keyPrototype().RegisterCodec(Codec)
	mappablePrototype().RegisterCodec(Codec)
	schema.RegisterCodec(Codec)
	Codec.Seal()

	return Codec
}
