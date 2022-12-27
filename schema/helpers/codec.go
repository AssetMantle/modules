package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

type Codec interface {
	client.TxConfig
	codec.Codec

	GetProtoCodec() *codec.ProtoCodec
	GetLegacyAmino() *codec.LegacyAmino
	InterfaceRegistry() types.InterfaceRegistry
	Initialize() Codec
}
