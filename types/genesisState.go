package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type GenesisState interface {
	Default() GenesisState
	Validate() error
	Initialize(sdkTypes.Context)
	Export(sdkTypes.Context) GenesisState
	Marshall() []byte
	Unmarshall([]byte) GenesisState
}
