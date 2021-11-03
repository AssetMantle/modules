/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*types.Data)(nil), nil)
	codec.RegisterInterface((*types.Fact)(nil), nil)
	codec.RegisterInterface((*types.Height)(nil), nil)
	codec.RegisterInterface((*types.ID)(nil), nil)
	codec.RegisterInterface((*types.MetaFact)(nil), nil)
	codec.RegisterInterface((*types.MetaProperties)(nil), nil)
	codec.RegisterInterface((*types.MetaProperty)(nil), nil)
	codec.RegisterInterface((*types.NFT)(nil), nil)
	codec.RegisterInterface((*types.Parameter)(nil), nil)
	codec.RegisterInterface((*types.Properties)(nil), nil)
	codec.RegisterInterface((*types.Property)(nil), nil)
	codec.RegisterInterface((*types.Signature)(nil), nil)
	codec.RegisterInterface((*types.Signatures)(nil), nil)

	codec.RegisterConcrete(&base.AccAddressData{}, "persistenceSpersistenceSDKDK.schema.types.base.AccAddressData", nil)
	codec.RegisterConcrete(&base.DecData{}, "persistenceSDK.schema.types.base.DecData", nil)
	codec.RegisterConcrete(&base.Fact{}, "persistenceSDK.schema.types.base.Fact", nil)
	codec.RegisterConcrete(&base.Height{}, "persistenceSDK.schema.types.base.Height", nil)
	codec.RegisterConcrete(&base.HeightData{}, "persistenceSDK.schema.types.base.HeightData", nil)
	codec.RegisterConcrete(&base.ID{}, "persistenceSDK.schema.types.base.ID", nil)
	codec.RegisterConcrete(&base.IDData{}, "persistenceSDK.schema.types.base.IDData", nil)
	codec.RegisterConcrete(&base.ListData{}, "persistenceSDK.schema.types.base.ListData", nil)
	codec.RegisterConcrete(&base.MetaFact{}, "persistenceSDK.schema.types.base.MetaFact", nil)
	codec.RegisterConcrete(&base.MetaProperties{}, "persistenceSDK.schema.types.base.MetaProperties", nil)
	codec.RegisterConcrete(&base.MetaProperty{}, "persistenceSDK.schema.types.base.MetaProperty", nil)
	codec.RegisterConcrete(&base.Properties{}, "persistenceSDK.schema.types.base.Properties", nil)
	codec.RegisterConcrete(&base.Property{}, "persistenceSDK.schema.types.base.Property", nil)
	codec.RegisterConcrete(&base.Signatures{}, "persistenceSDK.schema.types.base.Signatures", nil)
	codec.RegisterConcrete(&base.Signature{}, "persistenceSDK.schema.types.base.Signature", nil)
	codec.RegisterConcrete(&base.StringData{}, "persistenceSDK.schema.types.base.StringData", nil)
}

func RegisterInterfaces(registry codecTypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"persistenceSDK.schema.types.Data",
		(*types.Data)(nil),
		&base.AccAddressData{},
		&base.DecData{},
		&base.HeightData{},
		&base.IDData{},
		&base.ListData{},
		&base.StringData{},
	)

	registry.RegisterInterface("persistenceSDK.schema.types.ID", (*types.ID)(nil), &base.ID{})
	registry.RegisterInterface("persistenceSDK.schema.types.Height", (*types.Height)(nil), &base.Height{})
	registry.RegisterInterface("persistenceSDK.schema.types.Signature", (*types.Signature)(nil), &base.Signature{})
	registry.RegisterInterface("persistenceSDK.schema.types.Signatures", (*types.Signatures)(nil), &base.Signatures{})
	registry.RegisterInterface("persistenceSDK.schema.types.Fact", (*types.Fact)(nil), &base.Fact{})
	registry.RegisterInterface("persistenceSDK.schema.types.MetaFact", (*types.MetaFact)(nil), &base.MetaFact{})
	registry.RegisterInterface("persistenceSDK.schema.types.MetaProperties", (*types.MetaProperties)(nil), &base.MetaProperties{})
	registry.RegisterInterface("persistenceSDK.schema.types.Property", (*types.Property)(nil), &base.Property{})
	registry.RegisterInterface("persistenceSDK.schema.types.Properties", (*types.Properties)(nil), &base.Properties{})
}
