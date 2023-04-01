package base

import (
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	sdkCodecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	documentIDGetters "github.com/AssetMantle/modules/utilities/rest/idGetters/docs"
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
func (codec codec) Initialize(moduleBasicManager sdkModuleTypes.BasicManager) helpers.Codec {
	std.RegisterLegacyAminoCodec(codec.legacyAmino)
	std.RegisterInterfaces(codec.interfaceRegistry)
	schema.RegisterLegacyAminoCodec(codec.legacyAmino)
	documentIDGetters.RegisterLegacyAminoCodec(codec.legacyAmino)
	moduleBasicManager.RegisterLegacyAminoCodec(codec.legacyAmino)
	moduleBasicManager.RegisterInterfaces(codec.interfaceRegistry)
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
