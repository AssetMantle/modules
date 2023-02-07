package errors

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*error)(nil), nil)
	codec.RegisterInterface((*Error)(nil), nil)
}
