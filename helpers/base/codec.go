// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/helpers"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	sdkCodecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type codec struct {
	interfaceRegistry sdkCodecTypes.InterfaceRegistry
	sdkClient.TxConfig
	legacyAmino *sdkCodec.LegacyAmino
	*sdkCodec.ProtoCodec
}

var _ helpers.Codec = (*codec)(nil)

func (codec codec) GetProtoCodec() *sdkCodec.ProtoCodec {
	return codec.ProtoCodec
}
func (codec codec) GetLegacyAmino() *sdkCodec.LegacyAmino {
	return codec.legacyAmino
}
func (codec codec) InterfaceRegistry() sdkCodecTypes.InterfaceRegistry {
	return codec.interfaceRegistry
}
func (codec codec) Initialize(moduleManager helpers.ModuleManager) helpers.Codec {
	std.RegisterLegacyAminoCodec(codec.legacyAmino)
	std.RegisterInterfaces(codec.interfaceRegistry)
	moduleManager.RegisterLegacyAminoCodec(codec.legacyAmino)
	moduleManager.RegisterInterfaces(codec.interfaceRegistry)
	return codec
}

func CodecPrototype() helpers.Codec {
	codec := codec{}
	codec.interfaceRegistry = sdkCodecTypes.NewInterfaceRegistry()
	codec.ProtoCodec = sdkCodec.NewProtoCodec(codec.interfaceRegistry)
	codec.TxConfig = tx.NewTxConfig(codec, tx.DefaultSignModes)
	codec.legacyAmino = sdkCodec.NewLegacyAmino()
	return codec
}

// TestCodec
// Deprecated: Only for testing. Use CodecPrototype instead.
func TestCodec() helpers.Codec {
	return CodecPrototype().Initialize(NewModuleManager(auth.AppModule{}))
}
