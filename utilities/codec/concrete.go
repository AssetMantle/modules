/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package codec

import (
	"reflect"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/persistenceOne/persistenceSDK/constants"
)

func RegisterLegacyAminoXPRTConcrete(codec *codec.LegacyAmino, moduleName string, o interface{}) {
	codec.RegisterConcrete(o, constants.ProjectRoute+"/"+moduleName+"/"+reflect.TypeOf(o).PkgPath()+"/"+reflect.TypeOf(o).Name(), nil)
}

func RegisterXPRTImplementations(protoCodec codec.ProtoCodec, moduleName string, o interface{}) {

}
