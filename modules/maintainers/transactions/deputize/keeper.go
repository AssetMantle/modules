/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper           helpers.Mapper
	verifyAuxiliary  helpers.Auxiliary
	conformAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	maintainers := mapper.NewMaintainers(transactionKeeper.mapper, context)

	fromMaintainerID := mapper.NewMaintainerID(message.ClassificationID, message.FromID)
	fromMaintainer := maintainers.Fetch(fromMaintainerID).Get(fromMaintainerID)
	if fromMaintainer == nil || !fromMaintainer.CanAddMaintainer() {
		return newTransactionResponse(errors.NotAuthorized)
	}

	toMaintainerID := mapper.NewMaintainerID(message.ClassificationID, message.FromID)
	toMaintainer := maintainers.Fetch(toMaintainerID).Get(toMaintainerID)
	if toMaintainer != nil {
		return newTransactionResponse(errors.EntityAlreadyExists)
	}

	mutables := base.NewMutables(message.MaintainedTraits)
	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, nil, mutables)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	//TODO add classification maintainer ids
	maintainers = maintainers.Add(mapper.NewMaintainer(toMaintainerID, message.MaintainedTraits, message.AddMaintainer, message.RemoveMaintainer, message.MutateMaintainer))
	return newTransactionResponse(nil)
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
