/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package module

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

func RegisterLegacyAminoCodec(amino *codec.LegacyAmino, keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) *codec.AminoCodec {
	moduleCdc := codec.NewAminoCodec(amino)
	schema.RegisterLegacyAminoCodec(amino)
	cryptoCodec.RegisterCrypto(amino)
	amino.Seal()
	return moduleCdc
}
