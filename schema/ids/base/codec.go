// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, AssetID{})
	schema.RegisterModuleConcrete(codec, ClassificationID{})
	schema.RegisterModuleConcrete(codec, DataID{})
	schema.RegisterModuleConcrete(codec, HashID{})
	schema.RegisterModuleConcrete(codec, IdentityID{})
	schema.RegisterModuleConcrete(codec, MaintainerID{})
	schema.RegisterModuleConcrete(codec, OrderID{})
	schema.RegisterModuleConcrete(codec, OwnableID{})
	schema.RegisterModuleConcrete(codec, PropertyID{})
	schema.RegisterModuleConcrete(codec, SplitID{})
	schema.RegisterModuleConcrete(codec, StringID{})
}
