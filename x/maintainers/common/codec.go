// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"github.com/AssetMantle/modules/x/maintainers/key"
	"github.com/AssetMantle/modules/x/maintainers/mappable"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/utilities/module"
)

var LegacyAmino *codec.LegacyAmino

func init() {
	LegacyAmino = module.RegisterLegacyAminoCodec(key.Prototype, mappable.Prototype)
}
