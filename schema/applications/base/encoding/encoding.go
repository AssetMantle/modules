package encoding

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/persistenceOne/persistenceSDK/schema"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Marshaler
	TxConfig          client.TxConfig
	LegacyAmino       *codec.LegacyAmino
}

func MakeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		LegacyAmino:       amino,
	}
}

func (encodingConfig EncodingConfig) RegisterLegacyAminoCodec() {
	std.RegisterLegacyAminoCodec(encodingConfig.LegacyAmino)
	schema.RegisterLegacyAminoCodec(encodingConfig.LegacyAmino)

	// In the app add ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino) after calling this
}

func (encodingConfig EncodingConfig) RegisterInterfaces() {
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	schema.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	// In the app add ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry) after calling this
}
