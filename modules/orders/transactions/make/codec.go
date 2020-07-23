package make

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(Message{}, TransactionRoute+"/"+"message", nil)
	codec.RegisterConcrete(transactionRequest{}, TransactionRoute+"/"+"request", nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	utilities.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
