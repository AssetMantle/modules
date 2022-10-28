// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/documents"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified/base"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	maintainAuxiliary     helpers.Auxiliary
	renumerateAuxiliary   helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.AssetID))

	Mappable := assets.Get(key.NewKey(message.AssetID))
	if Mappable == nil {
		return newTransactionResponse(errorConstants.EntityNotFound)
	}
	asset := Mappable.(documents.Asset)

	if auxiliaryResponse := transactionKeeper.maintainAuxiliary.GetKeeper().Help(context, maintain.NewAuxiliaryRequest(asset.GetClassificationID(), message.FromID, base.NewMutables(baseLists.NewPropertyList(constants.SupplyProperty)))); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	metaProperties, err := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(asset.GetSupply())))
	if err != nil {
		return newTransactionResponse(err)
	}

	if supplyMetaProperty := metaProperties.GetMetaProperty(constants.SupplyProperty.GetID()); supplyMetaProperty != nil {
		value := supplyMetaProperty.GetData().(data.DecData).Get()
		if auxiliaryResponse := transactionKeeper.renumerateAuxiliary.GetKeeper().Help(context, renumerate.NewAuxiliaryRequest(message.FromID, message.AssetID, value)); !auxiliaryResponse.IsSuccessful() {
			return newTransactionResponse(auxiliaryResponse.GetError())
		}
	} else {
		return newTransactionResponse(errorConstants.MetaDataError)
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
