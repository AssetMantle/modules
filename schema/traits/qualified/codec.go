/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

const moduleName = "traits"

func RegisterLegacyCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, HasImmutables{})
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, moduleName, HasMutables{})
}
