package base

import (
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	sdkCodecTypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/AssetMantle/modules/schema/helpers"
)

type codec struct {
	sdkCodecTypes.InterfaceRegistry
	sdkCodec.Codec
	sdkClient.TxConfig
	legacyAmino *sdkCodec.LegacyAmino
	protoCodec  *sdkCodec.ProtoCodec
}

var _ helpers.Codec = (*codec)(nil)

func (codec codec) GetProtoCodec() *sdkCodec.ProtoCodec {
	return codec.protoCodec
}
func (codec codec) GetLegacyAmino() *sdkCodec.LegacyAmino {
	return codec.legacyAmino
}
func (codec codec) InitializeAndSeal() helpers.Codec {
	return codec
}
func (codec codec) UnpackAny(any *sdkCodecTypes.Any, iface interface{}) error {
	return codec.InterfaceRegistry.UnpackAny(any, iface)
}

func CodecPrototype() helpers.Codec {
	return &codec{}
}
