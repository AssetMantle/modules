package utilities

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Mapper interface {
	InitializeMapper(*codec.Codec, sdkTypes.StoreKey) Mapper
	RegisterCodec(*codec.Codec)
}
