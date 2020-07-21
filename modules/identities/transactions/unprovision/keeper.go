package unprovision

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type transactionKeeper struct {
	mapper utility.Mapper
}

var _ utility.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	identityID := message.IdentityID
	identities := mapper.NewIdentities(transactionKeeper.mapper, context).Fetch(identityID)
	identity := identities.Get(identityID)
	if identity == nil {
		return constants.EntityNotFound
	}
	if !identity.IsProvisioned(message.To) {
		return constants.EntityNotFound
	}
	if identity.IsUnprovisioned(message.To) {
		return constants.DeletionNotAllowed
	}
	identities.Mutate(identity.UnprovisionAddress(message.To))
	return nil
}

func initializeTransactionKeeper(mapper utility.Mapper, externalKeepers []interface{}) utility.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
