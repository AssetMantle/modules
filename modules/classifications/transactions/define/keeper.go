package define

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	var properties types.Properties
	for _, trait := range message.Traits.GetList() {
		properties.Add(trait.GetProperty())
	}
	mutables := base.NewMutables(properties, message.MaintainersID)
	immutables := base.NewImmutables(properties)
	classificationID := mapper.NewClassificationID(base.NewID(context.ChainID()), mutables.GetMaintainersID(), immutables.GetHashID())
	classifications := mapper.NewClassifications(transactionKeeper.mapper, context).Fetch(classificationID)
	if classifications.Get(classificationID) != nil {
		return constants.EntityAlreadyExists
	}
	classifications.Add(mapper.NewClassification(classificationID, message.Traits))
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, _ []interface{}) helpers.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
