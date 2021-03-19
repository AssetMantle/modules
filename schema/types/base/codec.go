/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

const moduleName = "types"

func RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, accAddressData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, decData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, fact{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, height{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, heightData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, id{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, idData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, metaFact{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, metaProperties{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, metaProperty{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, parameter{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, properties{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, property{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, signature{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, signatures{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, stringData{})
}
