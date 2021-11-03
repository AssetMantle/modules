/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

func RegisterLegacyCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*helpers.Message)(nil), nil)
	codec.RegisterInterface((*helpers.Key)(nil), nil)
	codec.RegisterInterface((*helpers.Mappable)(nil), nil)
	codec.RegisterInterface((*helpers.Parameters)(nil), nil)
}

func RegisterInterfaces(registry codecTypes.InterfaceRegistry) {
	registry.RegisterInterface("persistenceSDK.schema.helpers.Message", (*helpers.Message)(nil))
	registry.RegisterInterface("persistenceSDK.schema.helpers.Key", (*helpers.Key)(nil))
	registry.RegisterInterface("persistenceSDK.schema.helpers.Mappable", (*helpers.Mappable)(nil))
	registry.RegisterInterface("persistenceSDK.schema.helpers.Parameters", (*helpers.Parameters)(nil))
}
