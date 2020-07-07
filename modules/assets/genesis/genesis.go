package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

type genesisState struct{}

var _ types.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() types.GenesisState {
	return NewGenesisState()
}

func (genesisState genesisState) Validate() error { return nil }

func (genesisState genesisState) Initialize(sdkTypes.Context) {
}
func (genesisState genesisState) Export(sdkTypes.Context) types.GenesisState {
	return NewGenesisState()
}
func (genesisState genesisState) Marshall() []byte {
	return packageCodec.MustMarshalJSON(genesisState)
}
func (genesisState genesisState) Unmarshall(byte []byte) types.GenesisState {
	packageCodec.UnmarshalJSON(byte, &genesisState)
	return genesisState
}
func NewGenesisState() types.GenesisState {
	return genesisState{}
}
