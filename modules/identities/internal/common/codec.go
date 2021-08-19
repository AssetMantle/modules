/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package common

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/utilities/module"
)

var LegacyAminoCodec *codec.LegacyAmino

func init() {
	LegacyAminoCodec = module.RegisterLegacyAminoCodec(key.Prototype, mappable.Prototype)
}
