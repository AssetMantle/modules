package parameters

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*Parameter)(nil), nil)
}
