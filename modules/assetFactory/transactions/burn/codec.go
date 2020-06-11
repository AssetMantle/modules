package burn

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/types"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(Message{}, "assetFactory/burn", nil)
}

var packageCodec = codec.New()

func init() {
	RegisterCodec(packageCodec)
	types.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
