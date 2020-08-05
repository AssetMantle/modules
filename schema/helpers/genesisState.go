package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState interface {
	Default() GenesisState
	Validate() error
	Initialize(sdkTypes.Context, Mapper)
	Export(sdkTypes.Context, Mapper) GenesisState
	RegisterCodec(*codec.Codec)
	Marshall() []byte
	Unmarshall([]byte) GenesisState
}
