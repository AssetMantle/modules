/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package send

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	verifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}
	if message.Split.LTE(sdkTypes.ZeroDec()) {
		return newTransactionResponse(constants.NotAuthorized)
	}
	fromSplitID := mapper.NewSplitID(message.FromID, message.OwnableID)
	splits := mapper.NewSplits(transactionKeeper.mapper, context).Fetch(fromSplitID)
	fromSplit := splits.Get(fromSplitID)
	if fromSplit == nil {
		return newTransactionResponse(constants.EntityNotFound)
	}
	fromSplit = fromSplit.Send(message.Split).(mappables.Split)
	if fromSplit.GetSplit().LT(sdkTypes.ZeroDec()) {
		return newTransactionResponse(constants.NotAuthorized)
	} else if fromSplit.GetSplit().Equal(sdkTypes.ZeroDec()) {
		splits.Remove(fromSplit)
	} else {
		splits.Mutate(fromSplit)
	}

	toSplitID := mapper.NewSplitID(message.ToID, message.OwnableID)
	toSplit := splits.Fetch(toSplitID).Get(toSplitID)
	if toSplit == nil {
		splits.Add(mapper.NewSplit(toSplitID, message.Split))
	} else {
		splits.Mutate(toSplit.Receive(message.Split).(mappables.Split))
	}
	return newTransactionResponse(nil)
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
