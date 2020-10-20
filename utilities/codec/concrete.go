package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"reflect"
)

func RegisterXPRTConcrete(codec *codec.Codec, moduleName string, o interface{}) {
	codec.RegisterConcrete(o, constants.ProjectRoute+"/"+moduleName+"/"+reflect.TypeOf(o).PkgPath()+"/"+reflect.TypeOf(o).Name(), nil)
}
