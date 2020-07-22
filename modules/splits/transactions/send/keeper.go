package send

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type transactionKeeper struct {
	mapper utility.Mapper
}

var _ utility.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	//message := messageFromInterface(msg)
	//assetID := message.AssetID
	//assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	//asset := assets.Get(assetID)
	//if asset == nil {
	//	return constants.EntityNotFound
	//}
	//mutableProperties := asset.GetMutables().Get()
	//for _, property := range message.Properties.GetList() {
	//	if mutableProperties.Get(property.GetID()) == nil {
	//		return constants.EntityNotFound
	//	}
	//	mutableProperties = mutableProperties.Send(property)
	//}
	//asset = mapper.NewAsset(asset.GetID(), asset.GetBurn(), asset.GetLock(), asset.GetImmutables(), schema.NewMutables(mutableProperties, asset.GetMutables().GetMaintainersID()))
	//assets = assets.Send(asset)
	return nil
}

func initializeTransactionKeeper(mapper utility.Mapper, auxiliaryKeepers []interface{}) utility.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
