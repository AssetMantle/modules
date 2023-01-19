package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

type Codec interface {
	client.TxConfig
	codec.Codec

	GetProtoCodec() *codec.ProtoCodec
	GetLegacyAmino() *codec.LegacyAmino
	InterfaceRegistry() types.InterfaceRegistry
	Initialize(module.BasicManager) Codec
}
