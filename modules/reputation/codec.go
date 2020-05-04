package reputation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/reputation/transactions/feedback"
)

func RegisterCodec(codec *codec.Codec) {
	feedback.RegisterCodec(codec)
}
