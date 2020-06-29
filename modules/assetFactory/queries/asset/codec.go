package asset

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(queryRequest{}, fmt.Sprintf("/%v/%v/%v", constants.ModuleName, constants.AssetQuery, "request"), nil)
	codec.RegisterConcrete(queryResponse{}, fmt.Sprintf("/%v/%v/%v", constants.ModuleName, constants.AssetQuery, "response"), nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	mapper.RegisterCodec(packageCodec)
	types.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
