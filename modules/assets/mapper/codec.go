package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(Asset{}, ModuleRoute+"/"+"asset", nil)
	codec.RegisterConcrete(assetID{}, ModuleRoute+"/"+"assetID", nil)
	codec.RegisterConcrete(assets{}, ModuleRoute+"/"+"assets", nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
