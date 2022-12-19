// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, accAddressData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, booleanData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, decData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, heightData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, idData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, listData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, stringData{})
}
