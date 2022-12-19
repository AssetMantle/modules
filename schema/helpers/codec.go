package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

type Codec interface {
	types.InterfaceRegistry
	codec.Codec
	client.TxConfig
	GetLegacyAmino() *codec.LegacyAmino
	InitializeAndSeal() Codec
}
