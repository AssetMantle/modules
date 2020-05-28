package burn

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(message{}, "assetFactory/burn", nil)
}

var packageCodec = codec.New()

func init() {
	RegisterCodec(packageCodec)
	packageCodec.Seal()
}
