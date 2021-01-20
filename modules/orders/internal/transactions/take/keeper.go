/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package take

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	parameters          helpers.Parameters
	scrubAuxiliary      helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
	transferAuxiliary   helpers.Auxiliary
	verifyAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(errors.EntityNotFound)
	}

	orderID := message.OrderID
	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.New(orderID))
	order := orders.Get(key.New(orderID))

	if order == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}

	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.(mappables.Order).GetTakerID(), order.(mappables.Order).GetExchangeRate(), order.(mappables.Order).GetMakerOwnableSplit())))
	if Error != nil {
		newTransactionResponse(Error)
	}

	if takerIDProperty := metaProperties.GetMetaProperty(base.NewID(properties.TakerID)); takerIDProperty != nil {
		takerID, Error := takerIDProperty.GetMetaFact().GetData().AsID()
		if Error != nil {
			return newTransactionResponse(errors.MetaDataError)
		} else if !takerID.Equals(base.NewID("")) && !takerID.Equals(message.FromID) {
			return newTransactionResponse(errors.NotAuthorized)
		}
	}

	exchangeRateProperty := metaProperties.GetMetaProperty(base.NewID(properties.ExchangeRate))
	if exchangeRateProperty == nil {
		return newTransactionResponse(errors.MetaDataError)
	}

	exchangeRate, Error := exchangeRateProperty.GetMetaFact().GetData().AsDec()
	if Error != nil {
		return newTransactionResponse(errors.MetaDataError)
	}

	makerOwnableSplitProperty := metaProperties.GetMetaProperty(base.NewID(properties.MakerOwnableSplit))
	if makerOwnableSplitProperty == nil {
		return newTransactionResponse(errors.MetaDataError)
	}

	makerOwnableSplit, Error := makerOwnableSplitProperty.GetMetaFact().GetData().AsDec()
	if Error != nil {
		return newTransactionResponse(errors.MetaDataError)
	}

	sendTakerOwnableSplit := makerOwnableSplit.Mul(exchangeRate)
	sendMakerOwnableSplit := message.TakerOwnableSplit.Quo(exchangeRate)
	updatedMakerOwnableSplit := makerOwnableSplit.Sub(sendMakerOwnableSplit)

	switch {
	case updatedMakerOwnableSplit.Equal(sdkTypes.ZeroDec()):
		if message.TakerOwnableSplit.LT(sendTakerOwnableSplit) {
			return newTransactionResponse(errors.InsufficientBalance)
		}

		orders.Remove(order)
	case updatedMakerOwnableSplit.LT(sdkTypes.ZeroDec()):
		if message.TakerOwnableSplit.LT(sendTakerOwnableSplit) {
			return newTransactionResponse(errors.InsufficientBalance)
		}

		sendMakerOwnableSplit = makerOwnableSplit

		orders.Remove(order)
	default:
		sendTakerOwnableSplit = message.TakerOwnableSplit
		mutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewDecData(updatedMakerOwnableSplit))))))

		if Error != nil {
			return newTransactionResponse(Error)
		}

		order = mappable.NewOrder(orderID, order.(mappables.Order).GetImmutables(), order.(mappables.Order).GetMutables().Mutate(mutableProperties.GetList()...))
		orders.Mutate(order)
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, order.(mappables.Order).GetMakerID(), order.(mappables.Order).GetTakerOwnableID(), sendTakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), message.FromID, order.(mappables.Order).GetMakerOwnableID(), sendMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

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
		default:
			panic(errors.UninitializedUsage)
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
