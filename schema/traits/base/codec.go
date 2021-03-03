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

func RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, HasImmutables{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, HasMutables{})
}
