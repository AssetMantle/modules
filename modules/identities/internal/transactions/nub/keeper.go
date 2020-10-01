/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package nub

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	defineAuxiliary helpers.Auxiliary
	scrubAuxiliary  helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	immutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(base.NewID(properties.NubID), base.NewMetaFact(base.NewIDData(message.NubID))))))
	if Error != nil {
		return newTransactionResponse(Error)
	}
	immutables := base.NewImmutables(immutableProperties)
	classificationID, Error := define.GetClassificationIDFromResponse(transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(base.NewImmutables(base.NewProperties(base.NewProperty(base.NewID(properties.NubID), base.NewFact(base.NewIDData(base.NewID("")))))), base.NewMutables(base.NewProperties()))))
	if classificationID == nil && Error != nil {
		return newTransactionResponse(Error)
	}
	identityID := mapper.NewIdentityID(classificationID, immutables.GetHashID())
	identities := mapper.NewIdentities(transactionKeeper.mapper, context).Fetch(identityID)
	if identities.Get(identityID) != nil {
		return newTransactionResponse(errors.EntityAlreadyExists)
	}

	identities.Add(mapper.NewIdentity(identityID, []sdkTypes.AccAddress{message.From}, []sdkTypes.AccAddress{}, immutables, base.NewMutables(base.NewProperties())))
	return newTransactionResponse(nil)
}

func initializeTransactionKeeper(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case define.Auxiliary.GetName():
				transactionKeeper.defineAuxiliary = value
			case scrub.Auxiliary.GetName():
				transactionKeeper.scrubAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
