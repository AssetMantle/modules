package contract

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/contract/transactions/sign"
)

func RegisterCodec(codec *codec.Codec) {
	sign.RegisterCodec(codec)
}
