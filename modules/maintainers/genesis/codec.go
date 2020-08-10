/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/mapper"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func (genesisState) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(genesisState{}, mapper.ModuleRoute+"/"+"genesisState", nil)
}

var PackageCodec = codec.New()

func init() {
	GenesisState.RegisterCodec(PackageCodec)
	schema.RegisterCodec(PackageCodec)
	mapper.RegisterCodec(PackageCodec)
	PackageCodec.Seal()
}
