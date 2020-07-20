package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func (mapper mapper) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(asset{}, ModuleRoute+"/"+"order", nil)
	codec.RegisterConcrete(assetID{}, ModuleRoute+"/"+"assetID", nil)
	codec.RegisterConcrete(assets{}, ModuleRoute+"/"+"assets", nil)
}
