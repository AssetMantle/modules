package take

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/auxiliaries/swap"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionKeeper struct {
	mapper                    helpers.Mapper
	exchangesSwapAuxiliary    helpers.Auxiliary
	identitiesVerifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	orderID := message.OrderID
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(orderID)
	order := orders.Get(orderID)
	if order == nil {
		return constants.EntityNotFound
	}
	if Error := transactionKeeper.identitiesVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); Error != nil {
		return Error
	}

	if order.GetTakerID().String() != "" && message.FromID.Compare(order.GetTakerID()) != 0 {
		return constants.NotAuthorized
	}

	order = mapper.NewOrder(order.GetID(), order.GetBurn(), order.GetLock(), order.GetImmutables(),
		order.GetMakerID(), message.FromID, order.GetMakerAssetAmount(), order.GetMakerAssetData(), order.GetTakerAssetAmount(),
		order.GetTakerAssetData(), order.GetSalt())
	orders = orders.Mutate(order)
	if Error := transactionKeeper.exchangesSwapAuxiliary.GetKeeper().Help(context, swap.NewAuxiliaryRequest(order)); Error != nil {
		return Error
	}
	orders.Remove(order)
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case swap.Auxiliary.GetName():
				transactionKeeper.exchangesSwapAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.identitiesVerifyAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
