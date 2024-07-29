package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterInterface((*Mappable)(nil), nil)
	legacyAmino.RegisterInterface((*QueryResponse)(nil), nil)
	legacyAmino.RegisterInterface((*QueryRequest)(nil), nil)
	legacyAmino.RegisterInterface((*CommonTransactionRequest)(nil), nil)
	legacyAmino.RegisterInterface((*TransactionRequest)(nil), nil)
	legacyAmino.RegisterInterface((*Request)(nil), nil)
	legacyAmino.RegisterInterface((*Error)(nil), nil)
	legacyAmino.RegisterInterface((*error)(nil), nil)
}
