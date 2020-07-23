package utilities

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState interface {
	Default() GenesisState
	Validate() error
	Initialize(sdkTypes.Context)
	Export(sdkTypes.Context) GenesisState
	RegisterCodec(*codec.Codec)
	Marshall() []byte
	Unmarshall([]byte) GenesisState
}
