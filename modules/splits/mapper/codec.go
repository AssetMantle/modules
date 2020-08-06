/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	assetsMapper "github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(splits{}, ModuleRoute+"/"+"splits", nil)
	codec.RegisterConcrete(split{}, ModuleRoute+"/"+"split", nil)
	codec.RegisterConcrete(splitID{}, ModuleRoute+"/"+"splitID", nil)
}

var packageCodec = codec.New()

func init() {
	RegisterCodec(packageCodec)
	assetsMapper.Mapper.RegisterCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
