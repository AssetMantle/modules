package properties

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*MetaProperty)(nil), nil)
	codec.RegisterInterface((*Property)(nil), nil)
}
