// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/schema/documents"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified/base"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	authenticateAuxiliary helpers.Auxiliary
	defineAuxiliary       helpers.Auxiliary
	superAuxiliary        helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message.From, message.FromIdentity)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	mappable := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.FromIdentity)).Get(key.NewKey(message.FromIdentity))
	if mappable == nil {
		return newTransactionResponse(errorConstants.EntityNotFound)
	}
	identity := mappable.(documents.Identity)

	if !identity.(documents.Identity).IsProvisioned(message.From) {
		return newTransactionResponse(errorConstants.NotAuthorized)
	}

	immutables := base.NewImmutables(baseLists.NewPropertyList(append(message.ImmutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...))

	mutables := base.NewMutables(baseLists.NewPropertyList(append(append(message.MutableMetaProperties.GetList(), message.MutableProperties.GetList()...), constants.AuthenticationProperty)...))

	classificationID, err := define.GetClassificationIDFromResponse(transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(immutables, mutables)))
	if err != nil {
		return newTransactionResponse(err)
	}

	if auxiliaryResponse := transactionKeeper.superAuxiliary.GetKeeper().Help(context, super.NewAuxiliaryRequest(classificationID, message.FromIdentity, mutables)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case define.Auxiliary.GetName():
				transactionKeeper.defineAuxiliary = value
			case super.Auxiliary.GetName():
				transactionKeeper.superAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			}
		default:
			panic(errorConstants.UninitializedUsage)
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
