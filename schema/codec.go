/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	baseTypes "github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*error)(nil), nil)
	types.RegisterCodec(codec)
	baseTypes.RegisterCodec(codec)
	traits.RegisterCodec(codec)
	baseTraits.RegisterCodec(codec)
	mappables.RegisterCodec(codec)
	helpers.RegisterCodec(codec)
}
