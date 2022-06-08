// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

const moduleName = "ids"

func RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, moduleName, dataID{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, id{})
	codecUtilities.RegisterModuleConcrete(codec, moduleName, propertyID{})
}
