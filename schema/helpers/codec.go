package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Mappable)(nil), nil)
	codec.RegisterInterface((*QueryResponse)(nil), nil)
	codec.RegisterInterface((*QueryRequest)(nil), nil)
	codec.RegisterInterface((*TransactionRequest)(nil), nil)
}
