/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/initialize"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper                    helpers.Mapper
	identitiesVerifyAuxiliary helpers.Auxiliary
	metasInitializeAuxiliary  helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.identitiesVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	immutableTraits := base.NewImmutables(base.NewProperties(append(message.ImmutableMetaTraits.RemoveData().GetList(), message.ImmutableTraits.GetList()...)))
	mutableTraits := base.NewMutables(base.NewProperties(append(message.MutableMetaTraits.RemoveData().GetList(), message.MutableTraits.GetList()...)))

	classificationID := mapper.NewClassificationID(base.NewID(context.ChainID()), immutableTraits, mutableTraits)
	classifications := mapper.NewClassifications(transactionKeeper.mapper, context).Fetch(classificationID)
	if classifications.Get(classificationID) != nil {
		return newTransactionResponse(constants.EntityAlreadyExists)
	}

	if auxiliaryResponse := transactionKeeper.metasInitializeAuxiliary.GetKeeper().Help(context, initialize.NewAuxiliaryRequest(message.ImmutableMetaTraits, message.MutableMetaTraits)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	classifications = classifications.Add(mapper.NewClassification(classificationID, immutableTraits, mutableTraits))
	return newTransactionResponse(nil)
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case verify.Auxiliary.GetName():
				transactionKeeper.identitiesVerifyAuxiliary = value
			case initialize.Auxiliary.GetName():
				transactionKeeper.metasInitializeAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
