/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package execution

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	moduleConstants "github.com/persistenceOne/persistenceSDK/constants/modules"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/utilities"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func ExecuteImmediately(orderID types.ID, orders helpers.Collection, context sdkTypes.Context, supplementAuxiliary helpers.Auxiliary, transferAuxiliary helpers.Auxiliary, scrubAuxiliary helpers.Auxiliary) error {
	order := orders.Get(key.New(orderID)).(mappables.Order)

	orderExchangeRate, Error := order.(mappables.Order).GetExchangeRate().AsDec()
	if Error != nil {
		return Error
	}

	orderLeftOverMakerOwnableSplit, _, Error := utilities.GetOrderMakerOwnableSplitAndExpiry(supplementAuxiliary, context, order)
	if Error != nil {
		return Error
	}

	orderMutated := false

	// Assuming ExchangeRate is positive
	orderTakerOwnableID := order.GetTakerOwnableID()
	executableTakerOwnableID := order.GetMakerOwnableID()

	accumulator := func(mappableOrder helpers.Mappable) bool {
		executableOrder := mappableOrder.(mappables.Order)

		executableOrderExchangeRate, Error := executableOrder.GetExchangeRate().AsDec()
		if Error != nil {
			panic(Error)
		}

		executableOrderMakerOwnableSplit, expiry, Error := utilities.GetOrderMakerOwnableSplitAndExpiry(supplementAuxiliary, context, executableOrder)
		if Error != nil {
			panic(Error)
		}

		if expiry >= context.BlockHeight() {
			orderLeftOverTakerOwnableSplitDemanded := orderExchangeRate.Abs().Mul(orderLeftOverMakerOwnableSplit).MulTruncate(sdkTypes.SmallestDec())
			executableOrderTakerOwnableSplitDemanded := executableOrderExchangeRate.Abs().Mul(executableOrderMakerOwnableSplit).MulTruncate(sdkTypes.SmallestDec())

			if orderLeftOverMakerOwnableSplit.Mul(executableOrderMakerOwnableSplit).GTE(orderLeftOverTakerOwnableSplitDemanded.Mul(executableOrderTakerOwnableSplitDemanded)) {
				switch {
				case orderLeftOverMakerOwnableSplit.GT(executableOrderTakerOwnableSplitDemanded):
					// sending to buyer
					if auxiliaryResponse := transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), order.GetMakerID(), orderTakerOwnableID, executableOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
						panic(auxiliaryResponse.GetError())
					}
					// sending to executableOrder
					if auxiliaryResponse := transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), executableOrder.GetMakerID(), executableTakerOwnableID, executableOrderTakerOwnableSplitDemanded)); !auxiliaryResponse.IsSuccessful() {
						panic(auxiliaryResponse.GetError())
					}

					orderLeftOverMakerOwnableSplit = orderLeftOverMakerOwnableSplit.Sub(executableOrderTakerOwnableSplitDemanded)

					orders.Remove(executableOrder)
				case orderLeftOverMakerOwnableSplit.LT(executableOrderTakerOwnableSplitDemanded):
					// sending to buyer
					if auxiliaryResponse := transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), order.GetMakerID(), orderTakerOwnableID, orderLeftOverTakerOwnableSplitDemanded)); !auxiliaryResponse.IsSuccessful() {
						panic(auxiliaryResponse.GetError())
					}
					// sending to executableOrder
					if auxiliaryResponse := transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), executableOrder.GetMakerID(), executableTakerOwnableID, orderLeftOverMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
						panic(auxiliaryResponse.GetError())
					}

					mutableProperties, Error := scrub.GetPropertiesFromResponse(scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewDecData(executableOrderMakerOwnableSplit.Sub(orderLeftOverTakerOwnableSplitDemanded)))))))
					if Error != nil {
						panic(Error)
					}

					orders.Mutate(mappable.NewOrder(executableOrder.GetID(), executableOrder.GetImmutables(), executableOrder.GetMutables().Mutate(mutableProperties.GetList()...)))

					orderLeftOverMakerOwnableSplit = sdkTypes.ZeroDec()
				default:
					// case orderLeftOverMakerOwnableSplit.Equal(executableOrderTakerOwnableSplitDemanded):
					// sending to buyer
					if auxiliaryResponse := transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), order.GetMakerID(), orderTakerOwnableID, executableOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
						panic(auxiliaryResponse.GetError())
					}
					// sending to seller
					if auxiliaryResponse := transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), executableOrder.GetMakerID(), executableTakerOwnableID, orderLeftOverMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
						panic(auxiliaryResponse.GetError())
					}

					orders.Remove(executableOrder)

					orderLeftOverMakerOwnableSplit = sdkTypes.ZeroDec()
				}

				orderMutated = true
			}
		}

		if orderLeftOverMakerOwnableSplit.Equal(sdkTypes.ZeroDec()) {
			orders.Remove(order)

			return true
		}

		return false
	}

	if orderExchangeRate.IsPositive() {
		orders.Iterate(key.New(key.NewOrderID(order.GetClassificationID(), order.GetMakerOwnableID(), order.GetTakerOwnableID(), base.NewID(string(moduleConstants.NegativeExchangeRate)), base.NewID(""), base.NewID(""), base.NewImmutables(base.NewProperties()))), accumulator)
	} else {
		orderTakerOwnableID = order.GetMakerOwnableID()
		executableTakerOwnableID = order.GetTakerOwnableID()
		orders.Iterate(key.New(key.NewOrderID(order.GetClassificationID(), order.GetMakerOwnableID(), order.GetTakerOwnableID(), base.NewID(string(moduleConstants.PositiveExchangeRate)), base.NewID(""), base.NewID(""), base.NewImmutables(base.NewProperties()))), accumulator)
	}

	if !orderLeftOverMakerOwnableSplit.Equal(sdkTypes.ZeroDec()) && orderMutated {
		mutableProperties, Error := scrub.GetPropertiesFromResponse(scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewDecData(orderLeftOverMakerOwnableSplit))))))
		if Error != nil {
			return Error
		}

		orders.Mutate(mappable.NewOrder(orderID, order.GetImmutables(), order.GetMutables().Mutate(mutableProperties.GetList()...)))
	}

	return nil
}
