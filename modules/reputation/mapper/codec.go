package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(&baseInterNFT{}, "assetFactory/baseInterNFT", nil)
	codec.RegisterConcrete(&baseInterNFTAddress{}, "assetFactory/baseInterNFTAddress", nil)
}
