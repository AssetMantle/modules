package provision

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type transactionKeeper struct {
	mapper utilities.Mapper
}

var _ utilities.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	identityID := message.IdentityID
	identities := mapper.NewIdentities(transactionKeeper.mapper, context).Fetch(identityID)
	identity := identities.Get(identityID)
	if identity == nil {
		return constants.EntityNotFound
	}
	if identity.IsProvisioned(message.To) {
		return constants.EntityAlreadyExists
	}
	if identity.IsUnprovisioned(message.To) {
		return constants.DeletionNotAllowed
	}
	identities.Mutate(identity.ProvisionAddress(message.To))
	return nil
}

func initializeTransactionKeeper(mapper utilities.Mapper, _ []interface{}) utilities.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
