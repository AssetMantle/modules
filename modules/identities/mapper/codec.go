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
	codec.RegisterConcrete(identities{}, ModuleRoute+"/"+"identities", nil)
	codec.RegisterConcrete(identity{}, ModuleRoute+"/"+"identity", nil)
	codec.RegisterConcrete(identityID{}, ModuleRoute+"/"+"identityID", nil)
}

var packageCodec = codec.New()

func init() {
	RegisterCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
