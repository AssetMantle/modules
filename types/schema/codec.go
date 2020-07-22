package schema

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(fact{}, "xprt/fact", nil)
	codec.RegisterConcrete(height{}, "xprt/height", nil)
	codec.RegisterConcrete(id{}, "xprt/id", nil)
	codec.RegisterConcrete(immutables{}, "xprt/immutables", nil)
	codec.RegisterConcrete(mutables{}, "xprt/mutables", nil)
	codec.RegisterConcrete(properties{}, "xprt/properties", nil)
	codec.RegisterConcrete(property{}, "xprt/property", nil)
	codec.RegisterConcrete(signature{}, "xprt/signature", nil)
	codec.RegisterConcrete(signatures{}, "xprt/signatures", nil)
}
