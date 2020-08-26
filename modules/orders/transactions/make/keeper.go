/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	conformAuxiliary    helpers.Auxiliary
	mintAuxiliary       helpers.Auxiliary
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
	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, base.NewID(mapper.ModuleName), message.MakerOwnableID, message.MakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	immutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.ImmutableMetaProperties.GetMetaPropertyList()...)))
	if Error != nil {
		return newTransactionResponse(Error)
	}
	immutables := base.NewImmutables(base.NewProperties(append(immutableProperties.GetList(), message.ImmutableProperties.GetList()...)))

	orderID := mapper.NewOrderID(message.ClassificationID, message.MakerOwnableID, message.TakerOwnableID, message.FromID, immutables)
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(orderID)

	makerOwnableSplit := message.MakerOwnableSplit
	order := orders.Get(orderID)
	if order != nil {
		metaProperties, Error := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.GetMakerOwnableSplit())))
		if Error != nil {
			return newTransactionResponse(Error)
		}
		oldMakerOwnableSplitMetaProperty := metaProperties.GetMetaProperty(base.NewID(properties.MakerOwnableSplit))
		if oldMakerOwnableSplitMetaProperty == nil {
			return newTransactionResponse(errors.MetaDataError)
		} else {
			oldMakerOwnableSplit, Error := oldMakerOwnableSplitMetaProperty.GetMetaFact().GetData().AsDec()
			if Error != nil {
				newTransactionResponse(errors.MetaDataError)
			} else {
				makerOwnableSplit = oldMakerOwnableSplit.Add(makerOwnableSplit)
			}
		}
	}
	mutableMetaProperties := message.MutableMetaProperties.AddMetaProperty(base.NewMetaProperty(base.NewID(properties.Expiry), base.NewMetaFact(base.NewHeightData(base.NewHeight(message.ExpiresIn.Get()+context.BlockHeight())))))
	mutableMetaProperties = mutableMetaProperties.AddMetaProperty(base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewDecData(makerOwnableSplit))))
	mutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(mutableMetaProperties.GetMetaPropertyList()...)))
	if Error != nil {
		return newTransactionResponse(Error)
	}
	mutables := base.NewMutables(base.NewProperties(append(mutableProperties.GetList(), message.MutableProperties.GetList()...)))

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	if order != nil {
		orders.Mutate(mapper.NewOrder(orderID, immutables, order.GetMutables().Mutate(mutables.Get().GetList()...)))
	} else {
		orders.Add(mapper.NewOrder(orderID, immutables, mutables))
	}
	return newTransactionResponse(nil)
}

func initializeTransactionKeeper(mapper helpers.Mapper, externalKeepers []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case mint.Auxiliary.GetName():
				transactionKeeper.mintAuxiliary = value
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
