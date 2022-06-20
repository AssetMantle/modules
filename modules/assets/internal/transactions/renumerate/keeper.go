// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	constants2 "github.com/AssetMantle/modules/schema/helpers/constants"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties/constants"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	maintainAuxiliary   helpers.Auxiliary
	renumerateAuxiliary helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
	verifyAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(message.AssetID))

	asset := assets.Get(key.FromID(message.AssetID))
	if asset == nil {
		return newTransactionResponse(constants2.EntityNotFound)
	}

	if auxiliaryResponse := transactionKeeper.maintainAuxiliary.GetKeeper().Help(context, maintain.NewAuxiliaryRequest(asset.(mappables.Asset).GetClassificationID(), message.FromID, baseLists.NewPropertyList(constants.Supply))); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	metaProperties, err := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(asset.(mappables.Asset).GetSupply())))
	if err != nil {
		return newTransactionResponse(err)
	}

	if supplyMetaProperty := metaProperties.GetMetaProperty(constants.SupplyProperty); supplyMetaProperty != nil {
		value := supplyMetaProperty.GetData().(data.DecData).Get()
		if auxiliaryResponse := transactionKeeper.renumerateAuxiliary.GetKeeper().Help(context, renumerate.NewAuxiliaryRequest(message.FromID, message.AssetID, value)); !auxiliaryResponse.IsSuccessful() {
			return newTransactionResponse(auxiliaryResponse.GetError())
		}
	} else {
		return newTransactionResponse(constants2.MetaDataError)
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case maintain.Auxiliary.GetName():
				transactionKeeper.maintainAuxiliary = value
			case renumerate.Auxiliary.GetName():
				transactionKeeper.renumerateAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
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
