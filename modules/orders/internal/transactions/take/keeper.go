// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package take

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
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
		return newTransactionResponse(errorConstants.EntityNotFound)
	}

	orderID := message.OrderID
	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(orderID))
	order := orders.Get(key.NewKey(orderID)).(mappables.Order)

	if order == nil {
		return newTransactionResponse(errorConstants.EntityNotFound)
	}

	metaProperties, err := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.GetTakerID(), order.GetMakerOwnableSplit())))
	if err != nil {
		newTransactionResponse(err)
	}

	if takerIDProperty := metaProperties.GetMetaProperty(constants.TakerIDProperty); takerIDProperty != nil {
		takerID := takerIDProperty.GetData().(data.IDData).Get()
		if takerID.Compare(baseIDs.NewStringID("")) != 0 && takerID.Compare(message.FromID) != 0 {
			return newTransactionResponse(errorConstants.NotAuthorized)
		}
	}

	exchangeRate := order.GetExchangeRate().GetData().(data.DecData).Get()

	makerOwnableSplitProperty := metaProperties.GetMetaProperty(constants.MakerOwnableSplitProperty)
	if makerOwnableSplitProperty == nil {
		return newTransactionResponse(errorConstants.MetaDataError)
	}

	makerOwnableSplit := makerOwnableSplitProperty.GetData().(data.DecData).Get()

	makerReceiveTakerOwnableSplit := makerOwnableSplit.MulTruncate(exchangeRate).MulTruncate(sdkTypes.SmallestDec())
	takerReceiveMakerOwnableSplit := message.TakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(exchangeRate)

	switch updatedMakerOwnableSplit := makerOwnableSplit.Sub(takerReceiveMakerOwnableSplit); {
	case updatedMakerOwnableSplit.Equal(sdkTypes.ZeroDec()):
		if message.TakerOwnableSplit.LT(makerReceiveTakerOwnableSplit) {
			return newTransactionResponse(errorConstants.InsufficientBalance)
		}

		orders.Remove(order)
	case updatedMakerOwnableSplit.LT(sdkTypes.ZeroDec()):
		if message.TakerOwnableSplit.LT(makerReceiveTakerOwnableSplit) {
			return newTransactionResponse(errorConstants.InsufficientBalance)
		}

		takerReceiveMakerOwnableSplit = makerOwnableSplit

		orders.Remove(order)
	default:
		makerReceiveTakerOwnableSplit = message.TakerOwnableSplit
		mutableProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(updatedMakerOwnableSplit)))))

		if err != nil {
			return newTransactionResponse(err)
		}

		// TODO call order.mutate
		order = mappable.NewOrder(order.GetClassificationID(), order.GetImmutables(), baseQualified.NewMutables(order.GetImmutables().GetImmutablePropertyList().Mutate(mutableProperties.GetList()...)))
		orders.Mutate(order)
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, order.GetMakerID(), order.GetTakerOwnableID(), makerReceiveTakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(baseIDs.NewStringID(module.Name), message.FromID, order.GetMakerOwnableID(), takerReceiveMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
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
			panic(errorConstants.UninitializedUsage)
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
