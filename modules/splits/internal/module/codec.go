/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package module

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema"
)

var Codec *codec.Codec

func init() {
	Codec := codec.New()
	key.Prototype().RegisterCodec(Codec)
	mappable.Prototype().RegisterCodec(Codec)
	schema.RegisterCodec(Codec)
	Codec.Seal()
}
