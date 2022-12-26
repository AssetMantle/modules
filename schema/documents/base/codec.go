// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, asset{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, classification{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, Document{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, identity{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, maintainer{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, order{})
}
