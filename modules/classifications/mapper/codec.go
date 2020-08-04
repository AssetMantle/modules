package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(classifications{}, ModuleRoute+"/"+"classifications", nil)
	codec.RegisterConcrete(classification{}, ModuleRoute+"/"+"classification", nil)
	codec.RegisterConcrete(classificationID{}, ModuleRoute+"/"+"classificationID", nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
