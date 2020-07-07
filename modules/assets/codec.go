package assets

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assets/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mutate"
)

func registerCodec(codec *codec.Codec) {
	mapper.RegisterCodec(codec)

	asset.Query.RegisterCodec(codec)
	burn.Transaction.RegisterCodec(codec)
	mint.Transaction.RegisterCodec(codec)
	mutate.Transaction.RegisterCodec(codec)
}
