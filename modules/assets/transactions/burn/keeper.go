/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package burn

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	burnAuxiliary       helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
	verifyAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if Error := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); Error != nil {
		return Error
	}
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(message.AssetID)
	asset := assets.Get(message.AssetID)
	if asset == nil {
		return newTransactionResponse(constants.EntityNotFound)
	}
	supplementAuxiliaryResponse, Error := supplement.ValidateResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(asset.GetBurn())))
	if Error != nil {
		return newTransactionResponse(Error)
	}
	burnHeightMetaFact := supplementAuxiliaryResponse.MetaProperties.GetMetaProperty(base.NewID(constants.BurnProperty))
	if burnHeightMetaFact == nil {
		return newTransactionResponse(constants.EntityNotFound)
	}
	burnHeight, Error := burnHeightMetaFact.GetMetaFact().GetData().AsHeight()
	if Error != nil {
		return newTransactionResponse(Error)
	}
	if burnHeight.IsGreaterThan(base.NewHeight(context.BlockHeight())) {
		return newTransactionResponse(constants.NotAuthorized)
	}
	if Error := transactionKeeper.burnAuxiliary.GetKeeper().Help(context, burn.NewAuxiliaryRequest(message.FromID, message.AssetID, sdkTypes.SmallestDec())); Error != nil {
		return Error
	}
	assets.Remove(asset)
	return newTransactionResponse(nil)
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case burn.Auxiliary.GetName():
				transactionKeeper.burnAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
