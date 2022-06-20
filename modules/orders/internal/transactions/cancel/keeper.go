// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package cancel

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	constants2 "github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties/constants"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	parameters          helpers.Parameters
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

	if message.FromID.Compare(order.(mappables.Order).GetMakerID()) != 0 {
		return newTransactionResponse(constants2.NotAuthorized)
	}

	metaProperties, err := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.(mappables.Order).GetMakerOwnableSplit())))
	if err != nil {
		return newTransactionResponse(err)
	}

	makerOwnableSplitProperty := metaProperties.GetMetaProperty(constants.MakerOwnableSplitProperty)
	if makerOwnableSplitProperty == nil {
		return newTransactionResponse(constants2.MetaDataError)
	}

	makerOwnableSplit := makerOwnableSplitProperty.GetData().(data.DecData).Get()

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(baseIDs.NewID(module.Name), message.FromID, order.(mappables.Order).GetMakerOwnableID(), makerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	orders.Remove(order)

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, externalKeeper := range auxiliaries {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
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
