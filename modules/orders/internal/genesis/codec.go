/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func (genesis) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(genesis{}, mapper.ModuleRoute+"/"+"genesis", nil)
}

var packageCodec = codec.New()

func init() {
	Genesis.RegisterCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	mapper.Mapper.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
