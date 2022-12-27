package base

import (
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	sdkCodecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
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
func (codec codec) Initialize() helpers.Codec {
	std.RegisterLegacyAminoCodec(codec.legacyAmino)
	std.RegisterInterfaces(codec.interfaceRegistry)
	schema.RegisterLegacyAminoCodec(codec.legacyAmino)
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
