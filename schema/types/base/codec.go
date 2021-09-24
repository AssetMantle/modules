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

func RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, AccAddressData{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, ListData{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, DecData{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, Fact{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, Height{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, HeightData{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, ID{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, IDData{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, MetaFact{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, MetaProperties{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, MetaProperty{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, Properties{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, Property{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, Signature{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, Signatures{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, sortedDataList{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, StringData{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec,moduleName,  Data_StringData{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec,moduleName,  Data_IdData{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec,moduleName,  Data_ListData{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec,moduleName,  Data_HeightData{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec,moduleName,  Data_DecData{})

}
