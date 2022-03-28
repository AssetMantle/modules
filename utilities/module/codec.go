/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package module

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

func RegisterLegacyAminoCodec(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()
	keyPrototype().RegisterLegacyAminoCodec(Codec)
	mappablePrototype().RegisterLegacyAminoCodec(Codec)
	schema.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()
	return Codec
}
