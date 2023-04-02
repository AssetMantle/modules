// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyOwnableID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AssetID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, ClassificationID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, CoinID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, DataID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, HashID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, IdentityID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, MaintainerID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, OrderID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, PropertyID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, SplitID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, StringID{})

	legacyAmino.RegisterInterface((*isAnyID_Impl)(nil), nil)
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_AssetID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_ClassificationID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_CoinID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_DataID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_HashID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_IdentityID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_MaintainerID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_OrderID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_OwnableID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_PropertyID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_SplitID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyID_StringID{})

	legacyAmino.RegisterInterface((*isAnyOwnableID_Impl)(nil), nil)
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyOwnableID_AssetID{})
	codecUtilities.RegisterModuleConcrete(legacyAmino, AnyOwnableID_CoinID{})
}
