package mapper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
)

func (mapper mapper) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(asset{}, fmt.Sprintf("/%v/%v", constants.ModuleName, "asset"), nil)
	codec.RegisterConcrete(assetID{}, fmt.Sprintf("/%v/%v", constants.ModuleName, "assetID"), nil)
	codec.RegisterConcrete(assets{}, fmt.Sprintf("/%v/%v", constants.ModuleName, "assets"), nil)
}
