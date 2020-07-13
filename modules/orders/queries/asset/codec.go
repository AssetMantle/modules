package asset

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(queryRequest{}, fmt.Sprintf("/%v/%v/%v", constants.ModuleName, constants.AssetQuery, "request"), nil)
	codec.RegisterConcrete(queryResponse{}, fmt.Sprintf("/%v/%v/%v", constants.ModuleName, constants.AssetQuery, "response"), nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	types.RegisterCodec(packageCodec)
	mapper.Mapper.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
