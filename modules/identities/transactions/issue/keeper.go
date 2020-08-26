/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package issue

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper           helpers.Mapper
	scrubAuxiliary   helpers.Auxiliary
	conformAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	// TODO add maintainer and fromIdentity check
	immutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.ImmutableMetaProperties.GetMetaPropertyList()...)))
	if Error != nil {
		return newTransactionResponse(Error)
	}
	immutables := base.NewImmutables(base.NewProperties(append(immutableProperties.GetList(), message.ImmutableProperties.GetList()...)))

	identityID := mapper.NewIdentityID(message.ClassificationID, immutables.GetHashID())
	identities := mapper.NewIdentities(transactionKeeper.mapper, context).Fetch(identityID)
	if identities.Get(identityID) != nil {
		return newTransactionResponse(errors.EntityAlreadyExists)
	}

	mutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.MutableMetaProperties.GetMetaPropertyList()...)))
	if Error != nil {
		return newTransactionResponse(Error)
	}
	mutables := base.NewMutables(base.NewProperties(append(mutableProperties.GetList(), message.MutableProperties.GetList()...)))

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	identities.Add(mapper.NewIdentity(identityID, []sdkTypes.AccAddress{message.To}, []sdkTypes.AccAddress{}, immutables, mutables))
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
			case scrub.Auxiliary.GetName():
				transactionKeeper.scrubAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
