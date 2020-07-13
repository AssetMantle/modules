package mint

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(Message{}, fmt.Sprintf("/%v/%v/%v", constants.ModuleName, constants.MintTransaction, "Message"), nil)
	codec.RegisterConcrete(transactionRequest{}, fmt.Sprintf("/%v/%v/%v", constants.ModuleName, constants.MintTransaction, "request"), nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	types.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
