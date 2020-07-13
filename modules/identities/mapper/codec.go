package mapper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/identities/constants"
)

func (mapper mapper) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(identities{}, fmt.Sprintf("/%v/%v", constants.ModuleName, "identities"), nil)
	codec.RegisterConcrete(identity{}, fmt.Sprintf("/%v/%v", constants.ModuleName, "identity"), nil)
	codec.RegisterConcrete(identityID{}, fmt.Sprintf("/%v/%v", constants.ModuleName, "identityID"), nil)
}
