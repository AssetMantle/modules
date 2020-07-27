package burn

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(message.AssetID)
	asset := assets.Get(message.AssetID)
	if asset == nil {
		return constants.EntityNotFound
	}
	if !asset.CanBurn(base.NewHeight(context.BlockHeight())) {
		return constants.DeletionNotAllowed
	}
	assets.Remove(asset)
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, _ []interface{}) helpers.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
