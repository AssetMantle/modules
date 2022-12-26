// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, AccAddressData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, BooleanData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, DecData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, HeightData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, IDData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, ListData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, StringData{})

	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyDataList{})

}
