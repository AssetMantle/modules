/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/super"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	parameters      helpers.Parameters
	defineAuxiliary helpers.Auxiliary
	scrubAuxiliary  helpers.Auxiliary
	superAuxiliary  helpers.Auxiliary
	verifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	immutableMetaProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.ImmutableMetaProperties.GetList()...)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	immutableProperties := base.NewProperties(append(immutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...)

	mutableMetaProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.MutableMetaProperties.GetList()...)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	mutableProperties := base.NewProperties(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	classificationID, Error := define.GetClassificationIDFromResponse(transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(immutableProperties, mutableProperties)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	if auxiliaryResponse := transactionKeeper.superAuxiliary.GetKeeper().Help(context, super.NewAuxiliaryRequest(classificationID, message.FromID, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case define.Auxiliary.GetName():
				transactionKeeper.defineAuxiliary = value
			case scrub.Auxiliary.GetName():
				transactionKeeper.scrubAuxiliary = value
			case super.Auxiliary.GetName():
				transactionKeeper.superAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
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
