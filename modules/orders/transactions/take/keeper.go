/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package take

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	scrubAuxiliary      helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
	transferAuxiliary   helpers.Auxiliary
	verifyAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(constants.EntityNotFound)
	}
	orderID := message.OrderID
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(orderID)
	order := orders.Get(orderID)
	if order == nil {
		return newTransactionResponse(constants.EntityNotFound)
	}
	auxiliaryResponse, Error := supplement.ValidateResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.GetTakerID(), order.GetExchangeRate(), order.GetMakerOwnableSplit())))
	if Error != nil {
		newTransactionResponse(Error)
	}

	if takerIDProperty := auxiliaryResponse.MetaProperties.GetMetaProperty(base.NewID(constants.TakerIDProperty)); takerIDProperty != nil {
		takerID, Error := takerIDProperty.GetMetaFact().GetData().AsID()
		if Error == nil {
			if takerID.Compare(message.FromID) != 0 {
				newTransactionResponse(constants.NotAuthorized)
			}
		}
	}
	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, order.GetMakerID(), order.GetTakerOwnableID(), message.TakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	exchangeRateProperty := auxiliaryResponse.MetaProperties.GetMetaProperty(base.NewID(constants.ExchangeRateProperty))
	if exchangeRateProperty == nil {
		return newTransactionResponse(constants.MetaDataError)
	}
	exchangeRate, Error := exchangeRateProperty.GetMetaFact().GetData().AsDec()
	if Error != nil {
		return newTransactionResponse(constants.MetaDataError)
	}

	makerOwnableSplitProperty := auxiliaryResponse.MetaProperties.GetMetaProperty(base.NewID(constants.MakerOwnableSplitProperty))
	if makerOwnableSplitProperty == nil {
		return newTransactionResponse(constants.MetaDataError)
	}
	makerOwnableSplit, Error := makerOwnableSplitProperty.GetMetaFact().GetData().AsDec()
	if Error != nil {
		return newTransactionResponse(constants.MetaDataError)
	}

	sendMakerOwnableSplit := message.TakerOwnableSplit.Mul(exchangeRate)
	updatedMakerOwnableSplit := makerOwnableSplit.Sub(sendMakerOwnableSplit)
	if updatedMakerOwnableSplit.LT(sdkTypes.ZeroDec()) {
		return newTransactionResponse(constants.InsufficientBalance)
	} else if updatedMakerOwnableSplit.Equal(sdkTypes.ZeroDec()) {
		orders.Remove(order)
	} else {
		if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(mapper.ModuleName), message.FromID, order.GetMakerOwnableID(), makerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
			return newTransactionResponse(auxiliaryResponse.GetError())
		}
		scrubMetaMutablesAuxiliaryResponse, Error := scrub.ValidateResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(base.NewID(constants.MakerSplitProperty), base.NewMetaFact(base.NewDecData(makerOwnableSplit))))))
		if Error != nil {
			return newTransactionResponse(Error)
		}
		order = mapper.NewOrder(order.GetID(), order.GetImmutables(), order.GetMutables().Mutate(scrubMetaMutablesAuxiliaryResponse.Properties.GetList()...))
		orders = orders.Mutate(order)
	}
	return newTransactionResponse(nil)
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case scrub.Auxiliary.GetName():
				transactionKeeper.scrubAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
