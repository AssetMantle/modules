package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func RegisterCodec(codec *codec.Codec) {
	types.RegisterCodec(codec)
	base.RegisterCodec(codec)
	mappables.RegisterCodec(codec)
	mappers.RegisterCodec(codec)
	helpers.RegisterCodec(codec)
	codec.RegisterConcrete(sdkTypes.Coin{}, "cosmos-sdk/coin", nil)
	codec.RegisterInterface((*traits.Exchangeable)(nil), nil)

}
