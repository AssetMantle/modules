package mapper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
)

func (mapper mapper) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(Order{}, fmt.Sprintf("/%v/%v", constants.ModuleName, "Order"), nil)
	codec.RegisterConcrete(orderID{}, fmt.Sprintf("/%v/%v", constants.ModuleName, "orderID"), nil)
	codec.RegisterConcrete(orders{}, fmt.Sprintf("/%v/%v", constants.ModuleName, "orders"), nil)
}
