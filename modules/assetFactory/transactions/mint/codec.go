package mint

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/types"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(Message{}, "assetFactory/mint", nil)
	codec.RegisterConcrete(request{}, "assetFactory/request", nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	types.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
