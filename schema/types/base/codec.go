/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(decData{}, "xprt/decData", nil)
	codec.RegisterConcrete(fact{}, "xprt/fact", nil)
	codec.RegisterConcrete(height{}, "xprt/height", nil)
	codec.RegisterConcrete(heightData{}, "xprt/heightData", nil)
	codec.RegisterConcrete(id{}, "xprt/id", nil)
	codec.RegisterConcrete(idData{}, "xprt/idData", nil)
	codec.RegisterConcrete(immutables{}, "xprt/immutables", nil)
	codec.RegisterConcrete(metaFact{}, "xprt/metaFact", nil)
	codec.RegisterConcrete(metaProperties{}, "xprt/metaProperties", nil)
	codec.RegisterConcrete(metaProperty{}, "xprt/metaProperty", nil)
	codec.RegisterConcrete(mutables{}, "xprt/mutables", nil)
	codec.RegisterConcrete(properties{}, "xprt/properties", nil)
	codec.RegisterConcrete(property{}, "xprt/property", nil)
	codec.RegisterConcrete(signature{}, "xprt/signature", nil)
	codec.RegisterConcrete(signatures{}, "xprt/signatures", nil)
	codec.RegisterConcrete(stringData{}, "xprt/stringData", nil)
}
