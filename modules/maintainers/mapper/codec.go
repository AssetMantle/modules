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
	codec.RegisterConcrete(maintainers{}, ModuleRoute+"/"+"maintainers", nil)
	codec.RegisterConcrete(maintainer{}, ModuleRoute+"/"+"maintainer", nil)
	codec.RegisterConcrete(maintainerID{}, ModuleRoute+"/"+"maintainerID", nil)
}

var packageCodec = codec.New()

func init() {
	RegisterCodec(packageCodec)
	assetsMapper.Mapper.RegisterCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
