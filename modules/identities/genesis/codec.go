package genesis

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
)

func (genesisState) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(genesisState{}, mapper.ModuleRoute+"/"+"genesisState", nil)
}

var PackageCodec = codec.New()

func init() {
	GenesisState.RegisterCodec(PackageCodec)
	PackageCodec.Seal()
}
