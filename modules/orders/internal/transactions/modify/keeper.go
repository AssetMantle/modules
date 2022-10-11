// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

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
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
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

	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.OrderID))

	order := orders.Get(key.NewKey(message.OrderID)).(mappables.Order)
	if order == nil {
		return newTransactionResponse(errorConstants.EntityNotFound)
	Mappable := orders.Get(key.FromID(message.OrderID))
	if Mappable == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}
	order := Mappable.(mappables.Order)

	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.GetMakerOwnableSplit())))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	transferMakerOwnableSplit := sdkTypes.ZeroDec()

	if makerOwnableSplitProperty := metaProperties.GetMetaProperty(constants.MakerOwnableSplitProperty); makerOwnableSplitProperty != nil {
		oldMakerOwnableSplit := makerOwnableSplitProperty.GetData().(data.DecData).Get()

		transferMakerOwnableSplit = message.MakerOwnableSplit.Sub(oldMakerOwnableSplit)
	} else {
		return newTransactionResponse(errors.MetaDataError)
	}

	if order.GetMakerOwnableSplit().LT(sdkTypes.ZeroDec()) {
		if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, message.FromID, order.GetMakerOwnableID(), order.GetMakerOwnableSplit().Abs())); !auxiliaryResponse.IsSuccessful() {
			return newTransactionResponse(auxiliaryResponse.GetError())
		}
	} else if order.GetMakerOwnableSplit().GT(sdkTypes.ZeroDec()) {
		if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, module.ModuleIdentityID, order.GetMakerOwnableID(), order.GetMakerOwnableSplit())); !auxiliaryResponse.IsSuccessful() {
			return newTransactionResponse(auxiliaryResponse.GetError())
		}
	}

	mutableMetaProperties := message.MutableMetaProperties.Add(base.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(message.MakerOwnableSplit)))
	mutableMetaProperties = mutableMetaProperties.Add(base.NewMetaProperty(constants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(message.ExpiresIn.Get()+context.BlockHeight()))))
	mutableMetaProperties := message.MutableMetaProperties.Add(base.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(message.MakerOwnableSplit)))
	mutableMetaProperties = mutableMetaProperties.Add(base.NewMetaProperty(constants.ExpiryProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(message.ExpiresIn.Get()+context.BlockHeight()))))

	scrubbedMutableMetaProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(mutableMetaProperties.GetList()...)))
	if err != nil {
		return newTransactionResponse(err)
	}

	updatedMutables := order.GetMutables().Mutate(append(scrubbedMutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)
	updatedMutables := order.GetMutablePropertyList().Mutate(append(scrubbedMutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(order.GetClassificationID(), order.GetImmutables(), updatedMutables)); !auxiliaryResponse.IsSuccessful() {
	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(order.GetClassificationID(), order.GetImmutablePropertyList(), updatedMutables)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	orders.Remove(order)
	orders.Add(mappable.NewOrder(order.GetClassificationID(), order.GetImmutables(), updatedMutables))
	orders.Add(mappable.NewOrder(
		key.NewOrderID(
			order.GetClassificationID(),
			order.GetMakerOwnableID(),
			order.GetTakerOwnableID(),
			baseIDs.NewID(message.TakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(message.MakerOwnableSplit).String()),
			baseIDs.NewID(order.GetCreation().GetData().String()),
			order.GetMakerID(), order.GetImmutablePropertyList(),
		),
		order.GetImmutablePropertyList(),
		updatedMutables),
	)

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
