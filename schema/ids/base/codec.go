// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AssetID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, ClassificationID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, DataID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, HashID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, IdentityID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, MaintainerID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, OrderID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, OwnableID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, PropertyID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, SplitID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, StringID{})
}
