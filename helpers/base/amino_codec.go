package base

import (
	"github.com/AssetMantle/schema/codec/utilities"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *sdkCodec.LegacyAmino) {
	utilities.RegisterModuleConcrete(legacyAmino, commonTransactionRequest{})
}
