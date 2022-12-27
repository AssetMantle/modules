// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types"

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

func (transactionKeeper transactionKeeper) Transact(context types.Context, message helpers.Message) helpers.TransactionResponse {
	_, err := transactionKeeper.Handle(context.Context(), message.(*Message))
	return newTransactionResponse(err)
}
func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*Response, error) {
	fromAddress, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	mappable := transactionKeeper.mapper.NewCollection(types.UnwrapSDKContext(context)).Fetch(key.NewKey(message.FromID)).Get(key.NewKey(message.FromID))
	if mappable == nil {
		return nil, errorConstants.EntityNotFound
	}
	identity := mappable.(documents.Identity)

	if !identity.(documents.Identity).IsProvisioned(fromAddress) {
		return nil, errorConstants.NotAuthorized
	}

	immutables := base.NewImmutables(baseLists.NewPropertyList(append(message.ImmutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...))

	mutables := base.NewMutables(baseLists.NewPropertyList(append(append(message.MutableMetaProperties.GetList(), message.MutableProperties.GetList()...), constants.AuthenticationProperty)...))

	classificationID, err := define.GetClassificationIDFromResponse(transactionKeeper.defineAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), define.NewAuxiliaryRequest(immutables, mutables)))
	if err != nil {
		return nil, err
	}

	if auxiliaryResponse := transactionKeeper.superAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), super.NewAuxiliaryRequest(classificationID, message.FromID, mutables)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	return &Response{}, nil
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
