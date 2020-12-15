package transaction

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

func RegisterCodec(messagePrototype func() helpers.Message) *codec.Codec {
	Codec := codec.New()
	messagePrototype().RegisterCodec(Codec)
	schema.RegisterCodec(Codec)
	Codec.Seal()
	return Codec
}
