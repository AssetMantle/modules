package escrow

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/escrow/transactions/execute"
)

func RegisterCodec(codec *codec.Codec) {
	execute.RegisterCodec(codec)
}
