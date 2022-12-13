package properties

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*MetaProperty)(nil), nil)
	codec.RegisterInterface((*Property)(nil), nil)
}
