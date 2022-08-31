// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, assetID{})
	codecUtilities.RegisterModuleConcrete(codec, classificationID{})
	codecUtilities.RegisterModuleConcrete(codec, dataID{})
	codecUtilities.RegisterModuleConcrete(codec, hashID{})
	codecUtilities.RegisterModuleConcrete(codec, identityID{})
	codecUtilities.RegisterModuleConcrete(codec, maintainerID{})
	codecUtilities.RegisterModuleConcrete(codec, metaID{})
	codecUtilities.RegisterModuleConcrete(codec, orderID{})
	codecUtilities.RegisterModuleConcrete(codec, ownableID{})
	codecUtilities.RegisterModuleConcrete(codec, propertyID{})
	codecUtilities.RegisterModuleConcrete(codec, splitID{})
	codecUtilities.RegisterModuleConcrete(codec, stringID{})
}
