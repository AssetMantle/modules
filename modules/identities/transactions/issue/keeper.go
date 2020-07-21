package issue

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type transactionKeeper struct {
	mapper utility.Mapper
}

var _ utility.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	mutables := schema.NewMutables(message.Properties, message.MaintainersID)
	immutables := schema.NewImmutables(message.Properties)
	identityID := mapper.NewIdentityID(schema.NewID(context.ChainID()), mutables.GetMaintainersID(), message.ClassificationID, immutables.GetHashID())
	identities := mapper.NewIdentities(transactionKeeper.mapper, context).Fetch(identityID)
	if identities.Get(identityID) != nil {
		return constants.EntityAlreadyExists
	}
	identities.Add(mapper.NewIdentity(identityID, []sdkTypes.AccAddress{message.To}, []sdkTypes.AccAddress{}, immutables, mutables))
	return nil
}

func initializeTransactionKeeper(mapper utility.Mapper, externalKeepers []interface{}) utility.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
