/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package send

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	parameters      helpers.Parameters
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
	fromSplitID := key.NewSplitID(message.FromID, message.OwnableID)
	splits := transactionKeeper.mapper.NewCollection(context).Fetch(key.New(fromSplitID))
	fromSplit := splits.Get(key.New(fromSplitID))
	if fromSplit == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}
	fromSplit = fromSplit.(mappables.Split).Send(message.Split).(mappables.Split)
	if fromSplit.(mappables.Split).GetSplit().LT(sdkTypes.ZeroDec()) {
		return newTransactionResponse(errors.NotAuthorized)
	} else if fromSplit.(mappables.Split).GetSplit().Equal(sdkTypes.ZeroDec()) {
		splits.Remove(fromSplit)
	} else {
		splits.Mutate(fromSplit)
	}

	toSplitID := key.NewSplitID(message.ToID, message.OwnableID)
	toSplit := splits.Fetch(key.New(toSplitID)).Get(key.New(toSplitID))
	if toSplit == nil {
		splits.Add(mappable.NewSplit(toSplitID, message.Split))
	} else {
		splits.Mutate(toSplit.(mappables.Split).Receive(message.Split).(mappables.Split))
	}
	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters
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
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
