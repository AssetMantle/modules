// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameters            helpers.Parameters
	conformAuxiliary      helpers.Auxiliary
	scrubAuxiliary        helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, module.ModuleIdentityID, message.MakerOwnableID, message.MakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	immutableMetaProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.ImmutableMetaProperties.GetList()...)))
	if err != nil {
		return newTransactionResponse(err)
	}

	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(append(immutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...))
	exchangeRate := message.TakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(message.MakerOwnableSplit)
	orderID := baseIDs.NewOrderID(message.ClassificationID, message.MakerOwnableID, message.TakerOwnableID, exchangeRate, baseTypes.NewHeight(context.BlockHeight()), message.FromID, immutables)
	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(orderID))

	if order := orders.Get(key.NewKey(orderID)); order != nil {
		return newTransactionResponse(errorConstants.EntityAlreadyExists)
	}

	mutableMetaProperties := message.MutableMetaProperties.Add(baseProperties.NewMetaProperty(constants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(message.ExpiresIn.Get()+context.BlockHeight()))))
	mutableMetaProperties = mutableMetaProperties.Add(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(message.MakerOwnableSplit)))

	scrubbedMutableMetaProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(mutableMetaProperties.GetList()...)))
	if err != nil {
		return newTransactionResponse(err)
	}

	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(append(scrubbedMutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...))

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	order := base.NewOrder(message.ClassificationID, immutables, mutables)
	orders = orders.Add(mappable.NewMappable(order))

	// Order execution
	orderMutated := false
	orderLeftOverMakerOwnableSplit := message.MakerOwnableSplit

	accumulator := func(mappableOrder helpers.Mappable) bool {
		executableOrder := mappableOrder.(documents.Order)

		executableOrderTakerOwnableSplitDemanded := executableOrder.GetExchangeRate().MulTruncate(executableOrder.GetMakerOwnableSplit()).MulTruncate(sdkTypes.SmallestDec())

		if order.GetExchangeRate().MulTruncate(executableOrder.GetExchangeRate()).MulTruncate(sdkTypes.SmallestDec()).MulTruncate(sdkTypes.SmallestDec()).LTE(sdkTypes.OneDec()) {
			switch {
			case orderLeftOverMakerOwnableSplit.GT(executableOrderTakerOwnableSplitDemanded):
				// sending to buyer
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, order.GetMakerID(), order.GetTakerOwnableID(), executableOrder.GetMakerOwnableSplit())); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}
				// sending to executableOrder
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, executableOrder.GetMakerID(), order.GetMakerOwnableID(), executableOrderTakerOwnableSplitDemanded)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}

				orderLeftOverMakerOwnableSplit = orderLeftOverMakerOwnableSplit.Sub(executableOrderTakerOwnableSplitDemanded)

				orders.Remove(mappable.NewMappable(executableOrder))
			case orderLeftOverMakerOwnableSplit.LT(executableOrderTakerOwnableSplitDemanded):
				// sending to buyer
				sendToBuyer := orderLeftOverMakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(executableOrder.GetExchangeRate())
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, order.GetMakerID(), order.GetTakerOwnableID(), sendToBuyer)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}
				// sending to executableOrder
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, executableOrder.GetMakerID(), order.GetMakerOwnableID(), orderLeftOverMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}

				mutableProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(executableOrder.GetMakerOwnableSplit().Sub(sendToBuyer))))))
				if err != nil {
					panic(err)
				}

				orders.Mutate(mappable.NewMappable(base.NewOrder(executableOrder.GetClassificationID(), executableOrder.GetImmutables(), executableOrder.GetMutables().Mutate(mutableProperties.GetList()...))))

				orderLeftOverMakerOwnableSplit = sdkTypes.ZeroDec()
			default:
				// case orderLeftOverMakerOwnableSplit.Equal(executableOrderTakerOwnableSplitDemanded):
				// sending to buyer
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, order.GetMakerID(), order.GetTakerOwnableID(), executableOrder.GetMakerOwnableSplit())); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}
				// sending to seller
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, executableOrder.GetMakerID(), order.GetMakerOwnableID(), orderLeftOverMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}

				orders.Remove(mappable.NewMappable(executableOrder))

				orderLeftOverMakerOwnableSplit = sdkTypes.ZeroDec()
			}

			orderMutated = true
		}

		if orderLeftOverMakerOwnableSplit.Equal(sdkTypes.ZeroDec()) {
			orders.Remove(mappable.NewMappable(order))
			return true
		}

		return false
	}

	orders.Iterate(key.NewKey(baseIDs.NewOrderID(order.GetClassificationID(), order.GetTakerOwnableID(), order.GetMakerOwnableID(), sdkTypes.SmallestDec(), baseTypes.NewHeight(-1), baseIDs.NewIdentityID(nil, nil), baseQualified.NewImmutables(baseLists.NewPropertyList()))), accumulator)

	if !orderLeftOverMakerOwnableSplit.Equal(sdkTypes.ZeroDec()) && orderMutated {
		mutableProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(orderLeftOverMakerOwnableSplit)))))
		if err != nil {
			return newTransactionResponse(err)
		}

		orders.Mutate(mappable.NewMappable(base.NewOrder(order.GetClassificationID(), order.GetImmutables(), order.GetMutables().Mutate(mutableProperties.GetList()...))))
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
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			}
		default:
			panic(errorConstants.UninitializedUsage)
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
