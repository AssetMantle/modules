/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transaction

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

func RegisterLegacyAminoCodec(messagePrototype func() helpers.Message) *codec.LegacyAmino {
	Codec := &codec.LegacyAmino{}
	messagePrototype().RegisterLegacyAminoCodec(Codec)
	schema.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()

	return Codec
}
