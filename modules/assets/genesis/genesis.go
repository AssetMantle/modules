package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

//TODO define genesis state
type genesisState struct{}

var _ utilities.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() utilities.GenesisState {
	return newGenesisState()
}

func (genesisState genesisState) Validate() error { return nil }

func (genesisState genesisState) Initialize(sdkTypes.Context) {
}
func (genesisState genesisState) Export(sdkTypes.Context) utilities.GenesisState {
	return newGenesisState()
}
func (genesisState genesisState) Marshall() []byte {
	return packageCodec.MustMarshalJSON(genesisState)
}
func (genesisState genesisState) Unmarshall(byte []byte) utilities.GenesisState {
	if Error := packageCodec.UnmarshalJSON(byte, &genesisState); Error != nil {
		return nil
	}
	return genesisState
}

func newGenesisState() utilities.GenesisState {
	return genesisState{}
}

var GenesisState = newGenesisState()
