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

func RegisterCodec(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) *codec.Codec {
	Codec := codec.New()
	keyPrototype().RegisterCodec(Codec)
	mappablePrototype().RegisterCodec(Codec)
	schema.RegisterCodec(Codec)
	Codec.Seal()

	return Codec
}
