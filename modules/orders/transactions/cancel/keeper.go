package cancel

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/auxiliaries/custody"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper                    helpers.Mapper
	identitiesVerifyAuxiliary helpers.Auxiliary
	exchangesCustodyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(message.OrderID)
	order := orders.Get(message.OrderID)
	if order == nil {
		return constants.EntityNotFound
	}
	if Error := transactionKeeper.identitiesVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, order.GetMakerID())); Error != nil {
		return Error
	}
	if !order.CanBurn(base.NewHeight(context.BlockHeight())) {
		return constants.DeletionNotAllowed
	}

	if Error := transactionKeeper.exchangesCustodyAuxiliary.GetKeeper().Help(context, custody.NewAuxiliaryRequest(order, true)); Error != nil {
		return Error
	}
	orders.Remove(order)
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, externalKeepers []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case verify.Auxiliary.GetName():
				transactionKeeper.identitiesVerifyAuxiliary = value
			case custody.Auxiliary.GetName():
				transactionKeeper.exchangesCustodyAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
