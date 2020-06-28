package asset

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
	"strings"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(queryRequest{}, strings.Join([]string{Query.GetModuleName(), Query.GetName(), "request"}, "/"), nil)
	codec.RegisterConcrete(queryRequest{}, strings.Join([]string{Query.GetModuleName(), Query.GetName(), "response"}, "/"), nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	mapper.RegisterCodec(packageCodec)
	types.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
