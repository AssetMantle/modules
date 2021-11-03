/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
)

func RegisterLegacyCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*Chain)(nil), nil)
	codec.RegisterInterface((*Classification)(nil), nil)
	codec.RegisterInterface((*InterIdentity)(nil), nil)
	codec.RegisterInterface((*InterNFT)(nil), nil)
	codec.RegisterInterface((*Maintainer)(nil), nil)
	codec.RegisterInterface((*Meta)(nil), nil)
	codec.RegisterInterface((*Order)(nil), nil)
	codec.RegisterInterface((*Split)(nil), nil)
}

func RegisterInterfaces(registry codecTypes.InterfaceRegistry) {
	registry.RegisterInterface("persistenceSDK.schema.mappables.Chain", (*Chain)(nil))
	registry.RegisterInterface("persistenceSDK.schema.mappables.Classification", (*Classification)(nil))
	registry.RegisterInterface("persistenceSDK.schema.mappables.InterIdentity", (*InterIdentity)(nil))
	registry.RegisterInterface("persistenceSDK.schema.mappables.InterNFT", (*InterNFT)(nil))
	registry.RegisterInterface("persistenceSDK.schema.mappables.Maintainer", (*Maintainer)(nil))
	registry.RegisterInterface("persistenceSDK.schema.mappables.Meta", (*Meta)(nil))
	registry.RegisterInterface("persistenceSDK.schema.mappables.Order", (*Order)(nil))
	registry.RegisterInterface("persistenceSDK.schema.mappables.Split", (*Split)(nil))
}
