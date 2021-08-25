/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	baseTypes "github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*error)(nil), nil)
	types.RegisterLegacyAminoCodec(codec)
	baseTypes.RegisterLegacyAminoCodec(codec)
	traits.RegisterLegacyCodec(codec)
	baseTraits.RegisterLegacyCodec(codec)
	mappables.RegisterLegacyCodec(codec)
	helpers.RegisterLegacyCodec(codec)
}

func RegisterProtoCodec(registry codecTypes.InterfaceRegistry) {
	registry.RegisterImplementations((*types.ID)(nil), &baseTypes.ID{})
	registry.RegisterImplementations((*types.Height)(nil), &baseTypes.Height{})
	registry.RegisterImplementations((*types.Signature)(nil), &baseTypes.Signature{})
	registry.RegisterImplementations((*types.Signatures)(nil), &baseTypes.Signatures{})
	registry.RegisterImplementations((*types.Data)(nil), &baseTypes.AccAddressData{}, &baseTypes.HeightData{}, &baseTypes.IDData{}, &baseTypes.StringData{}, &baseTypes.DecData{})
	registry.RegisterImplementations((*types.Fact)(nil), &baseTypes.Fact{})
	registry.RegisterImplementations((*types.ListData)(nil), &baseTypes.ListData{})
	registry.RegisterImplementations((*types.MetaFact)(nil), &baseTypes.MetaFact{})
	registry.RegisterImplementations((*types.MetaProperties)(nil), &baseTypes.MetaProperties{})
	registry.RegisterImplementations((*types.Parameter)(nil), &baseTypes.Parameter{})
	registry.RegisterImplementations((*types.Property)(nil), &baseTypes.Property{})
	registry.RegisterImplementations((*types.Properties)(nil), &baseTypes.Properties{})
	registry.RegisterImplementations((*traits.HasMutables)(nil), &baseTraits.HasMutables{})
	registry.RegisterImplementations((*traits.HasImmutables)(nil), &baseTraits.HasImmutables{})
}
