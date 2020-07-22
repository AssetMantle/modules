package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func (mapper mapper) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(splits{}, ModuleRoute+"/"+"splits", nil)
	codec.RegisterConcrete(split{}, ModuleRoute+"/"+"split", nil)
	codec.RegisterConcrete(splitID{}, ModuleRoute+"/"+"splitID", nil)
}
