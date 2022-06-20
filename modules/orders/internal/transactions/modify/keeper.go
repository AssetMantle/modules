// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	constants2 "github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
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

	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(message.OrderID))

	order := orders.Get(key.FromID(message.OrderID))
	if order == nil {
		return newTransactionResponse(constants2.EntityNotFound)
	}

	metaProperties, err := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.(mappables.Order).GetMakerOwnableSplit())))
	if err != nil {
		return newTransactionResponse(err)
	}

	transferMakerOwnableSplit := sdkTypes.ZeroDec()

	if makerOwnableSplitProperty := metaProperties.GetMetaProperty(constants.MakerOwnableSplitProperty); makerOwnableSplitProperty != nil {
		oldMakerOwnableSplit := makerOwnableSplitProperty.GetData().(data.DecData).Get()

		transferMakerOwnableSplit = message.MakerOwnableSplit.Sub(oldMakerOwnableSplit)
	} else {
		return newTransactionResponse(constants2.MetaDataError)
	}

	if transferMakerOwnableSplit.LT(sdkTypes.ZeroDec()) {
		if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(baseIDs.NewID(module.Name), message.FromID, order.(mappables.Order).GetMakerOwnableID(), transferMakerOwnableSplit.Abs())); !auxiliaryResponse.IsSuccessful() {
			return newTransactionResponse(auxiliaryResponse.GetError())
		}
	} else if transferMakerOwnableSplit.GT(sdkTypes.ZeroDec()) {
		if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, baseIDs.NewID(module.Name), order.(mappables.Order).GetMakerOwnableID(), transferMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
			return newTransactionResponse(auxiliaryResponse.GetError())
		}
	}

	mutableMetaProperties := message.MutableMetaProperties.Add(base.NewMetaProperty(constants.MakerOwnableSplitProperty, baseData.NewDecData(message.MakerOwnableSplit)))
	mutableMetaProperties = mutableMetaProperties.Add(base.NewMetaProperty(constants.ExpiryProperty, baseData.NewHeightData(baseTypes.NewHeight(message.ExpiresIn.Get()+context.BlockHeight()))))

	scrubbedMutableMetaProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(mutableMetaProperties.GetList()...)))
	if err != nil {
		return newTransactionResponse(err)
	}

	updatedMutables := order.(mappables.Order).GetMutablePropertyList().Mutate(append(scrubbedMutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(order.(mappables.Order).GetClassificationID(), order.(mappables.Order).GetImmutablePropertyList(), updatedMutables)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	orders.Remove(order)
	orders.Add(mappable.NewOrder(
		key.NewOrderID(
			order.(mappables.Order).GetClassificationID(),
			order.(mappables.Order).GetMakerOwnableID(),
			order.(mappables.Order).GetTakerOwnableID(),
			baseIDs.NewID(message.TakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(message.MakerOwnableSplit).String()),
			baseIDs.NewID(order.(mappables.Order).GetCreation().GetData().String()),
			order.(mappables.Order).GetMakerID(), order.(mappables.Order).GetImmutablePropertyList(),
		),
		order.(mappables.Order).GetImmutablePropertyList(),
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
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			}
		default:
			panic(constants2.UninitializedUsage)
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
