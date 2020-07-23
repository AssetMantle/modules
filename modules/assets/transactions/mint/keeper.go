package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type transactionKeeper struct {
	mapper                    utilities.Mapper          `json:"mapper" valid:"required~Enter the mapper"`
	splitsMintAuxiliaryKeeper utilities.AuxiliaryKeeper `json:"splitsMintAuxiliaryKeeper" valid:"required~Enter the splitsMintAuxiliaryKeeper"`
}

var _ utilities.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	mutables := base.NewMutables(message.Properties, message.MaintainersID)
	immutables := base.NewImmutables(message.Properties)
	assetID := mapper.NewAssetID(base.NewID(context.ChainID()), mutables.GetMaintainersID(), message.ClassificationID, immutables.GetHashID())
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	if assets.Get(assetID) != nil {
		return constants.EntityAlreadyExists
	}
	assets.Add(mapper.NewAsset(assetID, message.Burn, message.Lock, immutables, mutables))
	return nil
}

func initializeTransactionKeeper(mapper utilities.Mapper, _ []interface{}) utilities.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	return transactionKeeper
}
