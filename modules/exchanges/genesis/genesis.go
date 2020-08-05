package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

//TODO define genesis state
type genesisState struct{}

var _ helpers.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() helpers.GenesisState {
	return newGenesisState()
}

func (genesisState genesisState) Validate() error { return nil }

func (genesisState genesisState) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {
}
func (genesisState genesisState) Export(sdkTypes.Context, helpers.Mapper) helpers.GenesisState {
	return newGenesisState()
}
func (genesisState genesisState) Marshall() []byte {
	return PackageCodec.MustMarshalJSON(genesisState)
}
func (genesisState genesisState) Unmarshall(byte []byte) helpers.GenesisState {
	if Error := PackageCodec.UnmarshalJSON(byte, &genesisState); Error != nil {
		return nil
	}
	return genesisState
}
func newGenesisState() helpers.GenesisState {
	return genesisState{}
}

var GenesisState = newGenesisState()