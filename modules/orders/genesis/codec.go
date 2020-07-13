package genesis

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
)

func (genesisState) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(genesisState{}, fmt.Sprintf("/%v/%v", constants.ModuleName, "genesisState"), nil)
}

var packageCodec = codec.New()

func init() {
	GenesisState.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
