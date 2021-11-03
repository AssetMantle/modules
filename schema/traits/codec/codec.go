/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/traits/base"
)

func RegisterLegacyCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*traits.HasImmutables)(nil), nil)
	codec.RegisterConcrete(&base.HasImmutables{}, "persistenceSDK.schema.traits.base.HasImmutables", nil)
	codec.RegisterInterface((*traits.HasMutables)(nil), nil)
	codec.RegisterConcrete(&base.HasMutables{}, "persistenceSDK.schema.traits.base.HasMutables", nil)
}

func RegisterInterfaces(registry codecTypes.InterfaceRegistry) {
	registry.RegisterInterface("persistenceSDK.schema.traits.HasMutables", (*traits.HasMutables)(nil), &base.HasMutables{})
	registry.RegisterInterface("persistenceSDK.schema.traits.HasImmutables", (*traits.HasImmutables)(nil), &base.HasImmutables{})
}
