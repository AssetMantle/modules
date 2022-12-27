// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"context"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
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
	"github.com/cosmos/cosmos-sdk/types"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameters            helpers.Parameters
	conformAuxiliary      helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context types.Context, message helpers.Message) helpers.TransactionResponse {
	_, err := transactionKeeper.Handle(context.Context(), message.(*Message))
	return newTransactionResponse(err)
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*Response, error) {

	fromAddress, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(message.FromID, module.ModuleIdentityID, message.MakerOwnableID, message.MakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	immutableMetaProperties := message.ImmutableMetaProperties.
		Add(baseProperties.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(message.TakerOwnableSplit.QuoTruncate(types.SmallestDec()).QuoTruncate(message.MakerOwnableSplit)))).
		Add(baseProperties.NewMetaProperty(constants.CreationHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(types.UnwrapSDKContext(context).BlockHeight())))).
		Add(baseProperties.NewMetaProperty(constants.MakerOwnableIDProperty.GetKey(), baseData.NewIDData(message.MakerOwnableID))).
		Add(baseProperties.NewMetaProperty(constants.TakerOwnableIDProperty.GetKey(), baseData.NewIDData(message.TakerOwnableID))).
		Add(baseProperties.NewMetaProperty(constants.MakerIDProperty.GetKey(), baseData.NewIDData(message.FromID))).
		Add(baseProperties.NewMetaProperty(constants.TakerIDProperty.GetKey(), baseData.NewIDData(message.TakerID)))

	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(append(immutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...))
	orderID := baseIDs.NewOrderID(message.ClassificationID, immutables)
	orders := transactionKeeper.mapper.NewCollection(types.UnwrapSDKContext(context)).Fetch(key.NewKey(orderID))

	if order := orders.Get(key.NewKey(orderID)); order != nil {
		return nil, errorConstants.EntityAlreadyExists
	}

	mutableMetaProperties := message.MutableMetaProperties.Add(baseProperties.NewMetaProperty(constants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(message.ExpiresIn.Get()+types.UnwrapSDKContext(context).BlockHeight()))))
	mutableMetaProperties = mutableMetaProperties.Add(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(message.MakerOwnableSplit)))

	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...))

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	order := base.NewOrder(message.ClassificationID, immutables, mutables)
	orders = orders.Add(mappable.NewMappable(order))

	// Order execution
	orderMutated := false
	orderLeftOverMakerOwnableSplit := message.MakerOwnableSplit

	accumulator := func(mappableOrder helpers.Mappable) bool {
		executableOrder := mappableOrder.(documents.Order)

		executableOrderTakerOwnableSplitDemanded := executableOrder.GetExchangeRate().MulTruncate(executableOrder.GetMakerOwnableSplit()).MulTruncate(types.SmallestDec())

		if order.GetExchangeRate().MulTruncate(executableOrder.GetExchangeRate()).MulTruncate(types.SmallestDec()).MulTruncate(types.SmallestDec()).LTE(types.OneDec()) {
			switch {
			case orderLeftOverMakerOwnableSplit.GT(executableOrderTakerOwnableSplitDemanded):
				// sending to buyer
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(module.ModuleIdentityID, order.GetMakerID(), order.GetTakerOwnableID(), executableOrder.GetMakerOwnableSplit())); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}
				// sending to executableOrder
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(module.ModuleIdentityID, executableOrder.GetMakerID(), order.GetMakerOwnableID(), executableOrderTakerOwnableSplitDemanded)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}

				orderLeftOverMakerOwnableSplit = orderLeftOverMakerOwnableSplit.Sub(executableOrderTakerOwnableSplitDemanded)

				orders.Remove(mappable.NewMappable(executableOrder))
			case orderLeftOverMakerOwnableSplit.LT(executableOrderTakerOwnableSplitDemanded):
				// sending to buyer
				sendToBuyer := orderLeftOverMakerOwnableSplit.QuoTruncate(types.SmallestDec()).QuoTruncate(executableOrder.GetExchangeRate())
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(module.ModuleIdentityID, order.GetMakerID(), order.GetTakerOwnableID(), sendToBuyer)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}
				// sending to executableOrder
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(module.ModuleIdentityID, executableOrder.GetMakerID(), order.GetMakerOwnableID(), orderLeftOverMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}

				mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(executableOrder.GetMakerOwnableSplit().Sub(sendToBuyer))))

				orders.Mutate(mappable.NewMappable(base.NewOrder(executableOrder.GetClassificationID(), executableOrder.GetImmutables(), executableOrder.GetMutables().Mutate(mutableProperties.GetList()...))))

				orderLeftOverMakerOwnableSplit = types.ZeroDec()
			default:
				// case orderLeftOverMakerOwnableSplit.Equal(executableOrderTakerOwnableSplitDemanded):
				// sending to buyer
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(module.ModuleIdentityID, order.GetMakerID(), order.GetTakerOwnableID(), executableOrder.GetMakerOwnableSplit())); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}
				// sending to seller
				if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(module.ModuleIdentityID, executableOrder.GetMakerID(), order.GetMakerOwnableID(), orderLeftOverMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
					panic(auxiliaryResponse.GetError())
				}

				orders.Remove(mappable.NewMappable(executableOrder))

				orderLeftOverMakerOwnableSplit = types.ZeroDec()
			}

			orderMutated = true
		}

		if orderLeftOverMakerOwnableSplit.Equal(types.ZeroDec()) {
			orders.Remove(mappable.NewMappable(order))
			return true
		}

		return false
	}

	orders.Iterate(mappable.NewMappable(order).GetKey(), accumulator)

	if !orderLeftOverMakerOwnableSplit.Equal(types.ZeroDec()) && orderMutated {
		mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(orderLeftOverMakerOwnableSplit)))

		orders.Mutate(mappable.NewMappable(base.NewOrder(order.GetClassificationID(), order.GetImmutables(), order.GetMutables().Mutate(mutableProperties.GetList()...))))
	}

	return &Response{}, nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, externalKeeper := range auxiliaries {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
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
