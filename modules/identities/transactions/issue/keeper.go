package issue

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type transactionKeeper struct {
	mapper utilities.Mapper `json:"mapper" valid:"required~Enter the Mapper"`
}

var _ utilities.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	mutables := base.NewMutables(message.Properties, message.MaintainersID)
	immutables := base.NewImmutables(message.Properties)
	identityID := mapper.NewIdentityID(base.NewID(context.ChainID()), mutables.GetMaintainersID(), message.ClassificationID, immutables.GetHashID())
	identities := mapper.NewIdentities(transactionKeeper.mapper, context).Fetch(identityID)
	if identities.Get(identityID) != nil {
		return constants.EntityAlreadyExists
	}
	identities.Add(mapper.NewIdentity(identityID, []sdkTypes.AccAddress{message.To}, []sdkTypes.AccAddress{}, immutables, mutables))
	return nil
}

func initializeTransactionKeeper(mapper utilities.Mapper, _ []interface{}) utilities.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
