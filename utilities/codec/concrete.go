// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package codec

import (
	"reflect"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/constants"
)

func RegisterModuleConcrete(codec *codec.Codec, moduleName string, o interface{}) {
	codec.RegisterConcrete(o, constants.ProjectRoute+"/"+moduleName+"/"+reflect.TypeOf(o).PkgPath()+"/"+reflect.TypeOf(o).Name(), nil)
}
