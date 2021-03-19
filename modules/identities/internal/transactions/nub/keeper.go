/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package nub

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mappable"
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

	nubIDProperty := base.NewMetaProperty(base.NewID(properties.NubID), base.NewMetaFact(base.NewIDData(message.NubID)))

	immutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(nubIDProperty)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	authenticationProperty := base.NewMetaProperty(base.NewID(properties.Authentication), base.NewMetaFact(base.NewAccAddressData(message.From)))

	mutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(authenticationProperty)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	classificationID, Error := define.GetClassificationIDFromResponse(transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(base.NewProperties(base.NewProperty(base.NewID(properties.NubID), base.NewFact(base.NewIDData(base.NewID(""))))), base.NewProperties(base.NewProperty(base.NewID(properties.Authentication), base.NewFact(base.NewAccAddressData(nil).ZeroValue()))))))
	if classificationID == nil && Error != nil {
		return newTransactionResponse(Error)
	}

	identityID := key.NewIdentityID(classificationID, immutableProperties)

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(identityID))
	if identities.Get(key.FromID(identityID)) != nil {
		return newTransactionResponse(errors.EntityAlreadyExists)
	}

	identities.Add(mappable.NewIdentity(identityID, immutableProperties, mutableProperties))

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case define.Auxiliary.GetName():
				transactionKeeper.defineAuxiliary = value
			case scrub.Auxiliary.GetName():
				transactionKeeper.scrubAuxiliary = value
			}
		default:
			panic(errors.UninitializedUsage)
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
