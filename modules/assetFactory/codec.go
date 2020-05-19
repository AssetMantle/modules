package assetFactory

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/lock"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/send"
)

func RegisterCodec(codec *codec.Codec) {
	mapper.RegisterCodec(codec)

	burn.RegisterCodec(codec)
	lock.RegisterCodec(codec)
	mint.RegisterCodec(codec)
	send.RegisterCodec(codec)
}
