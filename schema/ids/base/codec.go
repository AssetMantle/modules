// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, assetID{})
	schema.RegisterModuleConcrete(codec, classificationID{})
	schema.RegisterModuleConcrete(codec, dataID{})
	schema.RegisterModuleConcrete(codec, HashID{})
	schema.RegisterModuleConcrete(codec, identityID{})
	schema.RegisterModuleConcrete(codec, maintainerID{})
	schema.RegisterModuleConcrete(codec, orderID{})
	schema.RegisterModuleConcrete(codec, ownableID{})
	schema.RegisterModuleConcrete(codec, propertyID{})
	schema.RegisterModuleConcrete(codec, splitID{})
	schema.RegisterModuleConcrete(codec, StringID{})
}
