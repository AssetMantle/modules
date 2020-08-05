package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(identities{}, ModuleRoute+"/"+"identities", nil)
	codec.RegisterConcrete(Identity{}, ModuleRoute+"/"+"identity", nil)
	codec.RegisterConcrete(identityID{}, ModuleRoute+"/"+"identityID", nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
