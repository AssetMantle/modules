// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, assetID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, classificationID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, dataID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, hashID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, identityID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, maintainerID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, orderID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, ownableID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, propertyID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, splitID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, stringID{})
}
