package properties

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterInterface((*AnyProperty)(nil), nil)
	legacyAmino.RegisterInterface((*MetaProperty)(nil), nil)
	legacyAmino.RegisterInterface((*Property)(nil), nil)
}
