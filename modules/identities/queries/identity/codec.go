package identity

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/identities/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(queryRequest{}, fmt.Sprintf("/%v/%v/%v", constants.ModuleName, constants.IdentityQuery, "request"), nil)
	codec.RegisterConcrete(queryResponse{}, fmt.Sprintf("/%v/%v/%v", constants.ModuleName, constants.IdentityQuery, "response"), nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	types.RegisterCodec(packageCodec)
	mapper.Mapper.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
