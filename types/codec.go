package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Asset)(nil), nil)
	codec.RegisterInterface((*Address)(nil), nil)
	codec.RegisterInterface((*Share)(nil), nil)
}
