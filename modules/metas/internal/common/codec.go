// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/modules/metas/internal/mappable"
	codec2 "github.com/AssetMantle/modules/utilities"
)

var Codec *codec.LegacyAmino

func init() {
	Codec = codec2.MakeModuleCode(key.Prototype, mappable.Prototype)
}
