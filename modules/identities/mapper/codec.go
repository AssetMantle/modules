package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(identities{}, ModuleRoute+"/"+"identities", nil)
	codec.RegisterConcrete(identity{}, ModuleRoute+"/"+"identity", nil)
	codec.RegisterConcrete(identityID{}, ModuleRoute+"/"+"identityID", nil)
}
