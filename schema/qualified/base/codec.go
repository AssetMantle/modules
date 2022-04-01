// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

const moduleName = "traits"

func RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, Document{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, HasImmutables{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, HasMutables{})
}
