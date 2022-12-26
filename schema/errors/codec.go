package errors

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterInterface((*error)(nil), nil)
	legacyAmino.RegisterInterface((*Error)(nil), nil)
}
