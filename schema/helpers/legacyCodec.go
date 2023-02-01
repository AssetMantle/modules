package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterInterface((*Mappable)(nil), nil)
	legacyAmino.RegisterInterface((*Parameter)(nil), nil)
	legacyAmino.RegisterInterface((*QueryResponse)(nil), nil)
	legacyAmino.RegisterInterface((*QueryRequest)(nil), nil)
	legacyAmino.RegisterInterface((*TransactionRequest)(nil), nil)
}
