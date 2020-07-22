package burn

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type transactionKeeper struct {
	mapper utility.Mapper
}

var _ utility.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	splits := mapper.NewSplits(transactionKeeper.mapper, context).Fetch(message.SplitID)
	split := splits.Get(message.SplitID)
	if split == nil {
		return constants.EntityNotFound
	}
	//if !split.CanBurn(schema.NewHeight(context.BlockHeight())) {
	//	return constants.DeletionNotAllowed
	//}
	splits.Remove(split)
	return nil
}

func initializeTransactionKeeper(mapper utility.Mapper, auxiliaryKeepers []interface{}) utility.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
