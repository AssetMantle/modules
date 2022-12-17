package base

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

type globals struct {
	codec    *codec.Codec
	registry *types.InterfaceRegistry
}

func (g globals) GetCodec() codec.BinaryCodec {
	return *g.codec
}

func (g globals) GetRegistry() types.InterfaceRegistry {
	return *g.registry
}

var GlobalInstance *globals

func SetGlobals(codec *codec.Codec, registry *types.InterfaceRegistry) {
	GlobalInstance = &globals{
		codec:    codec,
		registry: registry,
	}
}
