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
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper           helpers.Mapper
	parameters       helpers.Parameters
	verifyAuxiliary  helpers.Auxiliary
	conformAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	maintainers := transactionKeeper.mapper.NewCollection(context)

	fromMaintainerID := key.NewMaintainerID(message.ClassificationID, message.FromID)
	fromMaintainer := maintainers.Fetch(key.New(fromMaintainerID)).Get(key.New(fromMaintainerID)).(mappables.Maintainer)
	if fromMaintainer == nil || !fromMaintainer.CanAddMaintainer() {
		return newTransactionResponse(errors.NotAuthorized)
	}

	toMaintainerID := key.NewMaintainerID(message.ClassificationID, message.ToID)
	toMaintainer := maintainers.Fetch(key.New(toMaintainerID)).Get(key.New(toMaintainerID))
	if toMaintainer != nil {
		return newTransactionResponse(errors.EntityAlreadyExists)
	}

	mutableTraits := base.NewMutables(message.MaintainedTraits)
	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, nil, mutableTraits)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	maintainers = maintainers.Add(mappable.NewMaintainer(toMaintainerID, mutableTraits, message.AddMaintainer, message.RemoveMaintainer, message.MutateMaintainer))
	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters
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
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
