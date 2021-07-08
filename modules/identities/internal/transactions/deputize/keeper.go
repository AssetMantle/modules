/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/deputize"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	parameters          helpers.Parameters
	deputizeAuxiliary   helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(message.FromID))
	identity := identities.Get(key.FromID(message.FromID))

	if identity == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}

	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(identity.(mappables.InterIdentity).GetAuthentication())))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	accAddressListData, ok := metaProperties.Get(base.NewID(properties.Authentication)).GetMetaFact().GetData().(types.ListData)
	if !ok {
		return newTransactionResponse(errors.EntityNotFound)
	}

	if !accAddressListData.IsPresent(base.NewAccAddressData(message.From)) {
		return newTransactionResponse(errors.NotAuthorized)

	}

	if auxiliaryResponse := transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {

		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	if auxiliaryResponse := transactionKeeper.deputizeAuxiliary.GetKeeper().Help(context, deputize.NewAuxiliaryRequest(message.FromID, message.ToID, message.ClassificationID, message.MaintainedProperties, message.AddMaintainer, message.RemoveMaintainer, message.MutateMaintainer)); !auxiliaryResponse.IsSuccessful() {

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
			case deputize.Auxiliary.GetName():
				transactionKeeper.deputizeAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
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
