// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/types/base"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	parameters          helpers.Parameters
	conformAuxiliary    helpers.Auxiliary
	scrubAuxiliary      helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
	transferAuxiliary   helpers.Auxiliary
	verifyAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, base.NewID(module.Name), message.MakerOwnableID, message.MakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	immutableMetaProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.ImmutableMetaProperties.GetList()...)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	immutableProperties := base.NewProperties(append(immutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...)
	exchangeRate := message.TakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(message.MakerOwnableSplit)
	orderID := key.NewOrderID(message.ClassificationID, message.MakerOwnableID, message.TakerOwnableID, base.NewID(exchangeRate.String()), base.NewID(strconv.FormatInt(context.BlockHeight(), 10)), message.FromID, immutableProperties)
	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(orderID))

	if order := orders.Get(key.FromID(orderID)); order != nil {
		return newTransactionResponse(errors.EntityAlreadyExists)
	}

	mutableMetaProperties := message.MutableMetaProperties.Add(base.NewMetaProperty(ids.ExpiryProperty, baseData.NewHeightData(base.NewHeight(message.ExpiresIn.Get()+context.BlockHeight()))))
	mutableMetaProperties = mutableMetaProperties.Add(base.NewMetaProperty(ids.MakerOwnableSplitProperty, baseData.NewDecData(message.MakerOwnableSplit)))

	scrubbedMutableMetaProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(mutableMetaProperties.GetList()...)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	mutableProperties := base.NewProperties(append(scrubbedMutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutableProperties, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	order := mappable.NewOrder(orderID, immutableProperties, mutableProperties)
	orders = orders.Add(order)

	// Order execution
	orderMutated := false
	orderLeftOverMakerOwnableSplit := message.MakerOwnableSplit

	orderExchangeRate, Error := order.GetExchangeRate().GetData().AsDec()
	if Error != nil {
		return newTransactionResponse(Error)
	}

	accumulator := func(mappableOrder helpers.Mappable) bool {
		executableOrder := mappableOrder.(mappables.Order)

		executableOrderExchangeRate, Error := executableOrder.GetExchangeRate().GetData().AsDec()
		if Error != nil {
			panic(Error)
		}

		executableOrderMetaProperties, Error := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(executableOrder.GetMakerOwnableSplit(), executableOrder.GetExpiry())))
		if Error != nil {
			panic(Error)
		}

		var executableOrderMakerOwnableSplit sdkTypes.Dec

		if makerOwnableSplitProperty := executableOrderMetaProperties.Get(ids.MakerOwnableSplitProperty); makerOwnableSplitProperty != nil {
			executableOrderMakerOwnableSplit, Error = makerOwnableSplitProperty.GetData().AsDec()
			if Error != nil {
				panic(Error)
			}
		} else {
			panic(errors.MetaDataError)
		}

		executableOrderTakerOwnableSplitDemanded := executableOrderExchangeRate.MulTruncate(executableOrderMakerOwnableSplit).MulTruncate(sdkTypes.SmallestDec())

		if orderExchangeRate.MulTruncate(executableOrderExchangeRate).MulTruncate(sdkTypes.SmallestDec()).MulTruncate(sdkTypes.SmallestDec()).LTE(sdkTypes.OneDec()) {
			switch {
			case orderLeftOverMakerOwnableSplit.GT(executableOrderTakerOwnableSplitDemanded):
				// sending to buyer
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), order.GetMakerID(), order.GetTakerOwnableID(), executableOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}
				// sending to executableOrder
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), executableOrder.GetMakerID(), order.GetMakerOwnableID(), executableOrderTakerOwnableSplitDemanded)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}

				orderLeftOverMakerOwnableSplit = orderLeftOverMakerOwnableSplit.Sub(executableOrderTakerOwnableSplitDemanded)

				orders.Remove(executableOrder)
			case orderLeftOverMakerOwnableSplit.LT(executableOrderTakerOwnableSplitDemanded):
				// sending to buyer
				sendToBuyer := orderLeftOverMakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(executableOrderExchangeRate)
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), order.GetMakerID(), order.GetTakerOwnableID(), sendToBuyer)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}
				// sending to executableOrder
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), executableOrder.GetMakerID(), order.GetMakerOwnableID(), orderLeftOverMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}

				mutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(ids.MakerOwnableSplitProperty, baseData.NewDecData(executableOrderMakerOwnableSplit.Sub(sendToBuyer))))))
				if Error != nil {
					panic(Error)
				}

				orders.Mutate(mappable.NewOrder(executableOrder.GetID(), executableOrder.GetImmutableProperties(), executableOrder.GetMutableProperties().Mutate(mutableProperties.GetList()...)))

				orderLeftOverMakerOwnableSplit = sdkTypes.ZeroDec()
			default:
				// case orderLeftOverMakerOwnableSplit.Equal(executableOrderTakerOwnableSplitDemanded):
				// sending to buyer
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), order.GetMakerID(), order.GetTakerOwnableID(), executableOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}
				// sending to seller
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), executableOrder.GetMakerID(), order.GetMakerOwnableID(), orderLeftOverMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}

				orders.Remove(executableOrder)

				orderLeftOverMakerOwnableSplit = sdkTypes.ZeroDec()
			}

			orderMutated = true
		}

		if orderLeftOverMakerOwnableSplit.Equal(sdkTypes.ZeroDec()) {
			orders.Remove(order)
			return true
		}

		return false
	}

	orders.Iterate(key.FromID(key.NewOrderID(order.GetClassificationID(), order.GetTakerOwnableID(), order.GetMakerOwnableID(), base.NewID(""), base.NewID(""), base.NewID(""), base.NewProperties())), accumulator)

	if !orderLeftOverMakerOwnableSplit.Equal(sdkTypes.ZeroDec()) && orderMutated {
		mutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(ids.MakerOwnableSplitProperty, baseData.NewDecData(orderLeftOverMakerOwnableSplit)))))
		if Error != nil {
			return newTransactionResponse(Error)
		}

		orders.Mutate(mappable.NewOrder(orderID, order.GetImmutableProperties(), order.GetMutableProperties().Mutate(mutableProperties.GetList()...)))
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, externalKeeper := range auxiliaries {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
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
