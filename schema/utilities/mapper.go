package utilities

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Mapper interface {
	Create(sdkTypes.Context, traits.Mappable)
	Read(sdkTypes.Context, types.ID) traits.Mappable
	Update(sdkTypes.Context, traits.Mappable)
	Delete(sdkTypes.Context, types.ID)
	Iterate(sdkTypes.Context, types.ID, func(traits.Mappable) bool)

	InitializeMapper(sdkTypes.StoreKey) Mapper
	RegisterCodec(*codec.Codec)
}
