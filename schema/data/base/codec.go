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
	codecUtilities.RegisterModuleConcrete(legacyAmino, NumberData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, StringData{})

	legacyAmino.RegisterInterface((*isAnyData_Impl)(nil), nil)
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyData_AccAddressData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyData_BooleanData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyData_DecData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyData_HeightData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyData_IDData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyData_ListData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyData_NumberData{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyData_StringData{})
}
