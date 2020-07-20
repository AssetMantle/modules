package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func (mapper mapper) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(order{}, ModuleRoute+"/"+"order", nil)
	codec.RegisterConcrete(orderID{}, ModuleRoute+"/"+"orderID", nil)
	codec.RegisterConcrete(orders{}, ModuleRoute+"/"+"orders", nil)
}
