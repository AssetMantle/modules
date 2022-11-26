package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(codec, parameter{})
}
