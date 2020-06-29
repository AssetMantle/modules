package assets

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assets/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mint"
)

func RegisterCodec(codec *codec.Codec) {
	mapper.RegisterCodec(codec)

	asset.Query.RegisterCodec(codec)
	mint.Transaction.RegisterCodec(codec)
}
