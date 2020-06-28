package mint

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/types"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(Message{}, fmt.Sprintf("/%v/%v/%v", Transaction.GetModuleName(), Transaction.GetName(), "message"), nil)
	codec.RegisterConcrete(transactionRequest{}, fmt.Sprintf("/%v/%v/%v", Transaction.GetModuleName(), Transaction.GetName(), "request"), nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	types.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
