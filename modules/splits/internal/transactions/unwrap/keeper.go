/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package unwrap

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	supplyKeeper    supply.Keeper
	verifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}
	if message.Split.LTE(sdkTypes.ZeroDec()) {
		return newTransactionResponse(errors.NotAuthorized)
	}
	splitID := mapper.NewSplitID(message.FromID, message.OwnableID)
	splits := mapper.NewSplits(transactionKeeper.mapper, context).Fetch(splitID)
	split := splits.Get(splitID)
	if split == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}
	split = split.Send(message.Split).(mappables.Split)
	if split.GetSplit().LT(sdkTypes.ZeroDec()) {
		return newTransactionResponse(errors.InsufficientBalance)
	} else if split.GetSplit().Equal(sdkTypes.ZeroDec()) {
		splits.Remove(split)
	} else {
		splits.Mutate(split)
	}
	if Error := transactionKeeper.supplyKeeper.SendCoinsFromModuleToAccount(context, mapper.ModuleName, message.From, sdkTypes.NewCoins(sdkTypes.NewCoin(message.OwnableID.String(), message.Split.TruncateInt()))); Error != nil {
		return newTransactionResponse(Error)
	}
	return newTransactionResponse(nil)
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case supply.Keeper:
			transactionKeeper.supplyKeeper = value
		case helpers.Auxiliary:
			switch value.GetName() {
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
