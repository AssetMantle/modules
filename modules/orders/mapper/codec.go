package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(Order{}, ModuleRoute+"/"+"order", nil)
	codec.RegisterConcrete(orderID{}, ModuleRoute+"/"+"orderID", nil)
	codec.RegisterConcrete(orders{}, ModuleRoute+"/"+"orders", nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
