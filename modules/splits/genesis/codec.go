package genesis

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
)

func (genesisState) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(genesisState{}, mapper.ModuleRoute+"/"+"genesisState", nil)
}

var packageCodec = codec.New()

func init() {
	GenesisState.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
