package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

//TODO define genesis state
type genesisState struct{}

var _ utility.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() utility.GenesisState {
	return newGenesisState()
}

func (genesisState genesisState) Validate() error { return nil }

func (genesisState genesisState) Initialize(sdkTypes.Context) {
}
func (genesisState genesisState) Export(sdkTypes.Context) utility.GenesisState {
	return newGenesisState()
}
func (genesisState genesisState) Marshall() []byte {
	return packageCodec.MustMarshalJSON(genesisState)
}
func (genesisState genesisState) Unmarshall(byte []byte) utility.GenesisState {
	packageCodec.UnmarshalJSON(byte, &genesisState)
	return genesisState
}

func newGenesisState() utility.GenesisState {
	return genesisState{}
}

var GenesisState = newGenesisState()
