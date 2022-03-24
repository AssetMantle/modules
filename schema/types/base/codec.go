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
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, booleanData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, listData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, decData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, height{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, heightData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, id{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, idData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, metaProperties{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, parameter{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, properties{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, property{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, signature{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, signatures{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, sortedDataList{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, stringData{})
}
