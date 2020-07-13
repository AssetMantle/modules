package issue

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper types.Mapper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	identities := mapper.NewIdentities(transactionKeeper.mapper, context).Fetch(message.IdentityID)
	identity := identities.Get(message.IdentityID)
	if identity == nil {
		return constants.EntityNotFound
	}
	identities.Remove(identity)
	return nil
}

func initializeTransactionKeeper(mapper types.Mapper) types.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
