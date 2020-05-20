package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Asset)(nil), nil)
	codec.RegisterInterface((*ID)(nil), nil)
	codec.RegisterInterface((*Share)(nil), nil)
}
