package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(&asset{}, "assetFactory/asset", nil)
	codec.RegisterConcrete(&assetID{}, "assetFactory/assetID", nil)
	codec.RegisterConcrete(&assets{}, "assetFactory/assets", nil)
}
