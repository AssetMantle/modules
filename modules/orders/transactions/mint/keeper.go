package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper types.Mapper
	//bankKeeper bankKeeper.Keeper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	return nil
}

func initializeTransactionKeeper(mapper types.Mapper) types.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
