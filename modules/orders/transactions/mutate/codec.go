package mutate

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(message{}, fmt.Sprintf("/%v/%v/%v", constants.ModuleName, constants.MutateTransaction, "message"), nil)
	codec.RegisterConcrete(transactionRequest{}, fmt.Sprintf("/%v/%v/%v", constants.ModuleName, constants.MutateTransaction, "request"), nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	types.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
