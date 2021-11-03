/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	helpersCodecs "github.com/persistenceOne/persistenceSDK/schema/helpers/codec"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	traitsCodecs "github.com/persistenceOne/persistenceSDK/schema/traits/codec"
	typesCodecs "github.com/persistenceOne/persistenceSDK/schema/types/codec"
)

func RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*error)(nil), nil)
	typesCodecs.RegisterLegacyAminoCodec(codec)
	traitsCodecs.RegisterLegacyCodec(codec)
	helpersCodecs.RegisterLegacyCodec(codec)
	mappables.RegisterLegacyCodec(codec)
}

func RegisterInterfaces(registry codecTypes.InterfaceRegistry) {
	typesCodecs.RegisterInterfaces(registry)
	traitsCodecs.RegisterInterfaces(registry)
	helpersCodecs.RegisterInterfaces(registry)
	mappables.RegisterInterfaces(registry)
}
