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

func RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, AccAddressData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, ListData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, DecData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, Fact{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, Height{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, HeightData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, ID{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, IDData{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, MetaFact{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, MetaProperties{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, MetaProperty{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, parameter{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, Properties{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, Property{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, Signature{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, Signatures{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, sortedDataList{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, StringData{})
}
