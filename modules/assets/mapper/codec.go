/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(asset{}, ModuleRoute+"/"+"asset", nil)
	codec.RegisterConcrete(assetID{}, ModuleRoute+"/"+"assetID", nil)
	codec.RegisterConcrete(assets{}, ModuleRoute+"/"+"assets", nil)
}

var packageCodec = codec.New()

func init() {
	RegisterCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
